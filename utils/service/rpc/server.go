package rpc

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"

	klog "github.com/kroma-network/kroma/utils/service/log"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	optls "github.com/kroma-network/kroma/utils/service/tls"
)

var wildcardHosts = []string{"*"}

type Server struct {
	endpoint       string
	apis           []rpc.API
	appVersion     string
	healthzHandler http.Handler
	corsHosts      []string
	vHosts         []string
	jwtSecret      []byte
	rpcPath        string
	healthzPath    string
	httpRecorder   kmetrics.HTTPRecorder
	httpServer     *http.Server
	log            log.Logger
	tls            *ServerTLSConfig
	middlewares    []Middleware
}

type ServerTLSConfig struct {
	Config    *tls.Config
	CLIConfig *optls.CLIConfig // paths to certificate and key files
}

type ServerOption func(b *Server)

type Middleware func(next http.Handler) http.Handler

func WithAPIs(apis []rpc.API) ServerOption {
	return func(b *Server) {
		b.apis = apis
	}
}

func WithHealthzHandler(hdlr http.Handler) ServerOption {
	return func(b *Server) {
		b.healthzHandler = hdlr
	}
}

func WithCORSHosts(hosts []string) ServerOption {
	return func(b *Server) {
		b.corsHosts = hosts
	}
}

func WithVHosts(hosts []string) ServerOption {
	return func(b *Server) {
		b.vHosts = hosts
	}
}

func WithJWTSecret(secret []byte) ServerOption {
	return func(b *Server) {
		b.jwtSecret = secret
	}
}

func WithRPCPath(path string) ServerOption {
	return func(b *Server) {
		b.rpcPath = path
	}
}

func WithHealthzPath(path string) ServerOption {
	return func(b *Server) {
		b.healthzPath = path
	}
}

func WithHTTPRecorder(recorder kmetrics.HTTPRecorder) ServerOption {
	return func(b *Server) {
		b.httpRecorder = recorder
	}
}

func WithLogger(lgr log.Logger) ServerOption {
	return func(b *Server) {
		b.log = lgr
	}
}

// WithTLSConfig configures TLS for the RPC server
// If this option is passed, the server will use ListenAndServeTLS
func WithTLSConfig(tls *ServerTLSConfig) ServerOption {
	return func(b *Server) {
		b.tls = tls
	}
}

// WithMiddleware adds an http.Handler to the rpc server handler stack
// The added middleware is invoked directly before the RPC callback
func WithMiddleware(middleware func(http.Handler) (hdlr http.Handler)) ServerOption {
	return func(b *Server) {
		b.middlewares = append(b.middlewares, middleware)
	}
}

func NewServer(host string, port int, appVersion string, opts ...ServerOption) *Server {
	endpoint := net.JoinHostPort(host, strconv.Itoa(port))
	bs := &Server{
		endpoint:       endpoint,
		appVersion:     appVersion,
		healthzHandler: defaultHealthzHandler(appVersion),
		corsHosts:      wildcardHosts,
		vHosts:         wildcardHosts,
		rpcPath:        "/",
		healthzPath:    "/healthz",
		httpRecorder:   kmetrics.NoopHTTPRecorder,
		httpServer: &http.Server{
			Addr: endpoint,
		},
		log: log.Root(),
	}
	for _, opt := range opts {
		opt(bs)
	}
	if bs.tls != nil {
		bs.httpServer.TLSConfig = bs.tls.Config
	}
	bs.AddAPI(rpc.API{
		Namespace: "health",
		Service: &healthzAPI{
			appVersion: appVersion,
		},
	})
	return bs
}

func (b *Server) Endpoint() string {
	return b.endpoint
}

func (b *Server) AddAPI(api rpc.API) {
	b.apis = append(b.apis, api)
}

func (b *Server) Start() error {
	srv := rpc.NewServer()
	if err := node.RegisterApis(b.apis, nil, srv); err != nil {
		return fmt.Errorf("error registering APIs: %w", err)
	}

	// rpc middleware
	var nodeHdlr http.Handler = srv
	for _, middleware := range b.middlewares {
		nodeHdlr = middleware(nodeHdlr)
	}
	nodeHdlr = node.NewHTTPHandlerStack(nodeHdlr, b.corsHosts, b.vHosts, b.jwtSecret)

	mux := http.NewServeMux()
	mux.Handle(b.rpcPath, nodeHdlr)
	mux.Handle(b.healthzPath, b.healthzHandler)

	// http middleware
	var handler http.Handler = mux
	handler = optls.NewPeerTLSMiddleware(handler)
	handler = kmetrics.NewHTTPRecordingMiddleware(b.httpRecorder, handler)
	handler = klog.NewLoggingMiddleware(b.log, handler)
	b.httpServer.Handler = handler

	errCh := make(chan error, 1)
	go func() {
		if b.tls != nil {
			if err := b.httpServer.ListenAndServeTLS("", ""); err != nil {
				errCh <- err
			}
		} else {
			if err := b.httpServer.ListenAndServe(); err != nil {
				errCh <- err
			}
		}
	}()

	// verify that the server comes up
	tick := time.NewTimer(10 * time.Millisecond)
	select {
	case err := <-errCh:
		return fmt.Errorf("http server failed: %w", err)
	case <-tick.C:
		return nil
	}
}

func (b *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return b.httpServer.Shutdown(ctx)
}

type HealthzResponse struct {
	Version string `json:"version"`
}

func defaultHealthzHandler(appVersion string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		_ = enc.Encode(&HealthzResponse{Version: appVersion})
	}
}

type healthzAPI struct {
	appVersion string
}

func (h *healthzAPI) Status() string {
	return h.appVersion
}
