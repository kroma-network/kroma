package watcher

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

type ContractWatcher struct {
	ctx     context.Context
	backend *ethclient.Client
	log     log.Logger
}

func NewContractWatcher(ctx context.Context, backend *ethclient.Client, l log.Logger) ContractWatcher {
	return ContractWatcher{
		ctx:     ctx,
		backend: backend,
		log:     l,
	}
}

func (cw ContractWatcher) WatchUpgraded(address common.Address, handlerFn func() error) error {
	proxy, err := bindings.NewProxy(address, cw.backend)
	if err != nil {
		return err
	}

	err = handlerFn()
	if err != nil {
		return err
	}

	opts := &bind.WatchOpts{Context: cw.ctx}
	upgradedCh := make(chan *bindings.ProxyUpgraded)
	sub := event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			cw.log.Warn("resubscribe contract upgraded event", "err", err)
		}
		return proxy.WatchUpgraded(opts, upgradedCh, nil)
	})

	go func() {
		for {
			select {
			case evt := <-upgradedCh:
				cw.log.Info("received contract upgraded event", "address", address)
				err := handlerFn()
				if err != nil {
					cw.log.Error("failed to update config", "err", err)
					<-time.After(time.Second)
					upgradedCh <- evt
					continue
				}
				cw.log.Info("config updated")
			case <-cw.ctx.Done():
				sub.Unsubscribe()
				close(upgradedCh)
				return
			}
		}
	}()

	return nil
}
