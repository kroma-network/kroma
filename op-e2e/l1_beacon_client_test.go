package op_e2e

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/fakebeacon"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
)

func TestGetVersion(t *testing.T) {
	InitParallel(t)

	l := testlog.Logger(t, log.LevelInfo)

	beaconApi := fakebeacon.NewBeacon(l, t.TempDir(), uint64(0), uint64(0))
	t.Cleanup(func() {
		_ = beaconApi.Close()
	})
	require.NoError(t, beaconApi.Start("127.0.0.1:0"))

	beaconCfg := sources.L1BeaconClientConfig{FetchAllSidecars: false}
	cl := sources.NewL1BeaconClient(sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(beaconApi.BeaconAddr(), l)), beaconCfg)

	version, err := cl.GetVersion(context.Background())
	require.NoError(t, err)
	require.Equal(t, "fakebeacon 1.2.3", version)
}
