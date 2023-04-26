package p2p_test

import (
	"testing"

	log "github.com/ethereum/go-ethereum/log"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	peer "github.com/libp2p/go-libp2p/core/peer"
	suite "github.com/stretchr/testify/suite"

	p2p "github.com/kroma-network/kroma/components/node/p2p"
	p2pMocks "github.com/kroma-network/kroma/components/node/p2p/mocks"
	"github.com/kroma-network/kroma/components/node/testlog"
)

// PeerScorerTestSuite tests peer parameterization.
type PeerScorerTestSuite struct {
	suite.Suite

	// mockConnGater *p2pMocks.ConnectionGater
	mockGater    *p2pMocks.PeerGater
	mockStore    *p2pMocks.Peerstore
	mockMetricer *p2pMocks.GossipMetricer
	bandScorer   *p2p.BandScoreThresholds
	logger       log.Logger
}

// SetupTest sets up the test suite.
func (testSuite *PeerScorerTestSuite) SetupTest() {
	testSuite.mockGater = &p2pMocks.PeerGater{}
	testSuite.mockStore = &p2pMocks.Peerstore{}
	testSuite.mockMetricer = &p2pMocks.GossipMetricer{}
	bandScorer, err := p2p.NewBandScorer("-40:graylist;0:friend;")
	testSuite.NoError(err)
	testSuite.bandScorer = bandScorer
	testSuite.logger = testlog.Logger(testSuite.T(), log.LvlError)
}

// TestPeerScorer runs the PeerScorerTestSuite.
func TestPeerScorer(t *testing.T) {
	suite.Run(t, new(PeerScorerTestSuite))
}

// TestScorer_OnConnect ensures we can call the OnConnect method on the peer scorer.
func (testSuite *PeerScorerTestSuite) TestScorer_OnConnect() {
	scorer := p2p.NewScorer(
		testSuite.mockGater,
		testSuite.mockStore,
		testSuite.mockMetricer,
		testSuite.bandScorer,
		testSuite.logger,
	)
	scorer.OnConnect()
}

// TestScorer_OnDisconnect ensures we can call the OnDisconnect method on the peer scorer.
func (testSuite *PeerScorerTestSuite) TestScorer_OnDisconnect() {
	scorer := p2p.NewScorer(
		testSuite.mockGater,
		testSuite.mockStore,
		testSuite.mockMetricer,
		testSuite.bandScorer,
		testSuite.logger,
	)
	scorer.OnDisconnect()
}

// TestScorer_SnapshotHook tests running the snapshot hook on the peer scorer.
func (testSuite *PeerScorerTestSuite) TestScorer_SnapshotHook() {
	scorer := p2p.NewScorer(
		testSuite.mockGater,
		testSuite.mockStore,
		testSuite.mockMetricer,
		testSuite.bandScorer,
		testSuite.logger,
	)
	inspectFn := scorer.SnapshotHook()

	// Mock the peer gater call
	testSuite.mockGater.On("Update", peer.ID("peer1"), float64(-100)).Return(nil).Once()

	// The metricer should then be called with the peer score band map
	testSuite.mockMetricer.On("SetPeerScores", map[string]float64{
		"friend":   0,
		"graylist": 1,
	}).Return(nil).Once()

	// Apply the snapshot
	snapshotMap := map[peer.ID]*pubsub.PeerScoreSnapshot{
		peer.ID("peer1"): {
			Score: -100,
		},
	}
	inspectFn(snapshotMap)

	// Change the peer score now to a different band
	testSuite.mockGater.On("Update", peer.ID("peer1"), float64(0)).Return(nil).Once()

	// The metricer should then be called with the peer score band map
	testSuite.mockMetricer.On("SetPeerScores", map[string]float64{
		"friend":   1,
		"graylist": 0,
	}).Return(nil).Once()

	// Apply the snapshot
	snapshotMap = map[peer.ID]*pubsub.PeerScoreSnapshot{
		peer.ID("peer1"): {
			Score: 0,
		},
	}
	inspectFn(snapshotMap)
}

// TestScorer_SnapshotHookBlocksPeer tests running the snapshot hook on the peer scorer with a peer score below the threshold.
// This implies that the peer should be blocked.
func (testSuite *PeerScorerTestSuite) TestScorer_SnapshotHookBlocksPeer() {
	scorer := p2p.NewScorer(
		testSuite.mockGater,
		testSuite.mockStore,
		testSuite.mockMetricer,
		testSuite.bandScorer,
		testSuite.logger,
	)
	inspectFn := scorer.SnapshotHook()

	// Mock the peer gater call
	testSuite.mockGater.On("Update", peer.ID("peer1"), float64(-101)).Return(nil)

	// The metricer should then be called with the peer score band map
	testSuite.mockMetricer.On("SetPeerScores", map[string]float64{
		"friend":   0,
		"graylist": 1,
	}).Return(nil)

	// Apply the snapshot
	snapshotMap := map[peer.ID]*pubsub.PeerScoreSnapshot{
		peer.ID("peer1"): {
			Score: -101,
		},
	}
	inspectFn(snapshotMap)
}
