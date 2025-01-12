// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package execution_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"testing"
	"time"

	"code.vegaprotocol.io/vega/core/collateral"
	"code.vegaprotocol.io/vega/core/epochtime"
	"code.vegaprotocol.io/vega/core/execution"
	"code.vegaprotocol.io/vega/core/integration/stubs"
	"code.vegaprotocol.io/vega/core/oracles"
	snp "code.vegaprotocol.io/vega/core/snapshot"
	"code.vegaprotocol.io/vega/core/stats"
	"code.vegaprotocol.io/vega/core/types"
	vgcontext "code.vegaprotocol.io/vega/libs/context"
	"code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/libs/num"
	"code.vegaprotocol.io/vega/libs/proto"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/paths"
	oraclespb "code.vegaprotocol.io/vega/protos/vega/oracles/v1"
	snapshot "code.vegaprotocol.io/vega/protos/vega/snapshot/v1"
	"github.com/stretchr/testify/require"
)

type snapshotTestData struct {
	engine         *execution.Engine
	oracleEngine   *oracles.Engine
	snapshotEngine *snp.Engine
	timeService    *stubs.TimeStub
}

// TestSnapshotOraclesTerminatingMarketFromSnapshot tests that market loaded from snapshot can be terminated with its oracle.
func TestSnapshotOraclesTerminatingMarketFromSnapshot(t *testing.T) {
	now := time.Now()
	exec := getEngine(t, now)
	mkt := newMarket("MarketID", "0xDEADBEEF")
	err := exec.engine.SubmitMarket(context.Background(), mkt, "")
	require.NoError(t, err)

	state, _, _ := exec.engine.GetState("")

	exec2 := getEngine(t, now)
	snap := &snapshot.Payload{}
	proto.Unmarshal(state, snap)
	_, _ = exec2.engine.LoadState(context.Background(), types.PayloadFromProto(snap))

	state2, _, _ := exec2.engine.GetState("")

	err = exec.engine.StartOpeningAuction(context.Background(), mkt.ID)
	require.NoError(t, err)
	mktState, err := exec.engine.GetMarketState("MarketID")
	require.NoError(t, err)
	require.Equal(t, types.MarketStateActive, mktState)

	err = exec2.engine.StartOpeningAuction(context.Background(), mkt.ID)
	require.NoError(t, err)
	mktState, err = exec2.engine.GetMarketState("MarketID")
	require.NoError(t, err)
	require.Equal(t, types.MarketStateActive, mktState)

	exec.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
		PubKeys: []string{"0xDEADBEEF"},
		Data:    map[string]string{"trading.terminated": "true"},
	})

	exec2.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
		PubKeys: []string{"0xDEADBEEF"},
		Data:    map[string]string{"trading.terminated": "true"},
	})

	marketState1, _ := exec.engine.GetMarketState("MarketID")
	marketState2, _ := exec2.engine.GetMarketState("MarketID")

	require.Equal(t, marketState1, marketState2)
	require.Equal(t, types.MarketStateTradingTerminated, marketState1)
	require.Equal(t, types.MarketStateTradingTerminated, marketState2)

	exec.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
		PubKeys: []string{"0xDEADBEEF"},
		Data:    map[string]string{"prices.ETH.value": "100"},
	})

	exec2.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
		PubKeys: []string{"0xDEADBEEF"},
		Data:    map[string]string{"prices.ETH.value": "100"},
	})

	marketState1, _ = exec.engine.GetMarketState("MarketID")
	marketState2, _ = exec2.engine.GetMarketState("MarketID")
	require.Equal(t, marketState1, marketState2)
	require.Equal(t, types.MarketStateSettled, marketState1)
	require.Equal(t, types.MarketStateSettled, marketState2)

	require.True(t, bytes.Equal(state, state2))
}

// TestSnapshotOraclesTerminatingMarketFromSnapshotAfterSettlementPrice sets up a market that gets the settlement price first.
// Then a snapshot is taken and another node is restored from this snapshot. Finally trading termination data is received and both markets
// are expected to get settled.
func TestSnapshotOraclesTerminatingMarketFromSnapshotAfterSettlementPrice(t *testing.T) {
	now := time.Now()
	exec := getEngine(t, now)
	mkt := newMarket("MarketID", "0xDEADBEEF")
	err := exec.engine.SubmitMarket(context.Background(), mkt, "")
	require.NoError(t, err)

	// settlement price arrives first
	exec.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
		PubKeys: []string{"0xDEADBEEF"},
		Data:    map[string]string{"prices.ETH.value": "100"},
	})

	// take a snapshot
	state, _, _ := exec.engine.GetState("")

	// load from the snapshot
	exec2 := getEngine(t, now)
	snap := &snapshot.Payload{}
	proto.Unmarshal(state, snap)
	_, _ = exec2.engine.LoadState(context.Background(), types.PayloadFromProto(snap))

	// take a snapshot on the loaded engine
	state2, _, _ := exec2.engine.GetState("")
	require.True(t, bytes.Equal(state, state2))

	err = exec.engine.StartOpeningAuction(context.Background(), mkt.ID)
	require.NoError(t, err)
	mktState, err := exec.engine.GetMarketState("MarketID")
	require.NoError(t, err)
	require.Equal(t, types.MarketStateActive, mktState)

	err = exec2.engine.StartOpeningAuction(context.Background(), mkt.ID)
	require.NoError(t, err)
	mktState, err = exec2.engine.GetMarketState("MarketID")
	require.NoError(t, err)
	require.Equal(t, types.MarketStateActive, mktState)

	// terminate the market to lead to settlement
	exec.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
		PubKeys: []string{"0xDEADBEEF"},
		Data:    map[string]string{"trading.terminated": "true"},
	})

	exec2.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
		PubKeys: []string{"0xDEADBEEF"},
		Data:    map[string]string{"trading.terminated": "true"},
	})

	// take snapshot for both engines, and verify they're both settled
	marketState1, _ := exec.engine.GetMarketState("MarketID")
	marketState2, _ := exec2.engine.GetMarketState("MarketID")
	require.Equal(t, marketState1, marketState2)
	require.Equal(t, types.MarketStateSettled, marketState1)
	require.Equal(t, types.MarketStateSettled, marketState2)
}

// TestLoadTerminatedMarketFromSnapshot terminates markets, loads them using the snapshot engine and then settles them successfully.
func TestLoadTerminatedMarketFromSnapshot(t *testing.T) {
	now := time.Now()
	exec := getEngine(t, now)
	defer exec.snapshotEngine.Close()
	ctx := vgcontext.WithTraceID(vgcontext.WithBlockHeight(context.Background(), 100), "0xDEADBEEF")
	ctx = vgcontext.WithChainID(ctx, "chainid")

	pubKeys := []string{"0xDEADBEEF", "0xDEADBEFF", "0xDEADBFFF"}
	marketIDs := []string{"market1", "market2", "market3"}

	// submit and terminate all markets
	for i := 0; i < 3; i++ {
		mkt := newMarket(marketIDs[i], pubKeys[i])
		err := exec.engine.SubmitMarket(ctx, mkt, "")
		require.NoError(t, err)

		// verify markets are terminated
		marketState, err := exec.engine.GetMarketState(marketIDs[i])
		require.NoError(t, err)
		require.Equal(t, types.MarketStateProposed, marketState)

		err = exec.engine.StartOpeningAuction(context.Background(), mkt.ID)
		require.NoError(t, err)
		marketState, err = exec.engine.GetMarketState(marketIDs[i])
		require.NoError(t, err)
		require.Equal(t, marketState, types.MarketStateActive)

		// terminate all markets
		exec.oracleEngine.BroadcastData(ctx, oracles.OracleData{
			PubKeys: []string{pubKeys[i]},
			Data:    map[string]string{"trading.terminated": "true"},
		})

		marketState, err = exec.engine.GetMarketState(marketIDs[i])
		require.NoError(t, err)
		require.Equal(t, types.MarketStateTradingTerminated, marketState)
	}

	// we now have 3 terminated markets in the execution engine
	// let's take a snapshot
	_, err := exec.snapshotEngine.Snapshot(ctx)
	require.NoError(t, err)
	snaps, err := exec.snapshotEngine.List()
	require.NoError(t, err)
	snap1 := snaps[0]

	// now let's start from this snapshot
	exec2 := getEngine(t, now)
	defer exec2.snapshotEngine.Close()
	exec2.snapshotEngine.ReceiveSnapshot(snap1)
	exec2.snapshotEngine.ApplySnapshot(ctx)
	exec2.snapshotEngine.CheckLoaded()

	// progress time to trigger any side effect on time ticks
	exec.timeService.SetTime(now.Add(2 * time.Second))
	exec2.timeService.SetTime(now.Add(2 * time.Second))

	// finally take a snapshot of both and compare them
	snp, err := exec.snapshotEngine.Snapshot(ctx)
	require.NoError(t, err)
	snp2, err := exec2.snapshotEngine.Snapshot(ctx)
	require.NoError(t, err)
	require.True(t, bytes.Equal(snp, snp2))

	// settle the markets
	for i := 0; i < 3; i++ {
		exec.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
			PubKeys: []string{pubKeys[i]},
			Data:    map[string]string{"prices.ETH.value": "100"},
		})
		exec2.oracleEngine.BroadcastData(context.Background(), oracles.OracleData{
			PubKeys: []string{pubKeys[i]},
			Data:    map[string]string{"prices.ETH.value": "100"},
		})

		marketState1, _ := exec.engine.GetMarketState(marketIDs[i])
		marketState2, _ := exec2.engine.GetMarketState(marketIDs[i])
		require.Equal(t, marketState1.String(), marketState2.String())
		require.Equal(t, types.MarketStateSettled, marketState1)
		require.Equal(t, types.MarketStateSettled, marketState2)

		// finally take a snapshot of both and compare them
		snp, err := exec.snapshotEngine.Snapshot(ctx)
		require.NoError(t, err)
		snp2, err := exec2.snapshotEngine.Snapshot(ctx)
		require.NoError(t, err)
		require.True(t, bytes.Equal(snp, snp2))
	}
}

func newMarket(ID, pubKey string) *types.Market {
	return &types.Market{
		ID: ID, // ID will be generated
		PriceMonitoringSettings: &types.PriceMonitoringSettings{
			Parameters: &types.PriceMonitoringParameters{
				Triggers: []*types.PriceMonitoringTrigger{
					{
						Horizon:          1000,
						HorizonDec:       num.DecimalFromFloat(1000.0),
						Probability:      num.DecimalFromFloat(0.3),
						AuctionExtension: 10000,
					},
				},
			},
		},
		LiquidityMonitoringParameters: &types.LiquidityMonitoringParameters{
			TargetStakeParameters: &types.TargetStakeParameters{
				TimeWindow:    100,
				ScalingFactor: num.DecimalFromFloat(1.0),
			},
			TriggeringRatio:  num.DecimalFromFloat(0.9),
			AuctionExtension: 10000,
		},
		Fees: &types.Fees{
			Factors: &types.FeeFactors{
				MakerFee:          num.DecimalFromFloat(0.1),
				InfrastructureFee: num.DecimalFromFloat(0.1),
				LiquidityFee:      num.DecimalFromFloat(0.1),
			},
		},
		TradableInstrument: &types.TradableInstrument{
			MarginCalculator: &types.MarginCalculator{
				ScalingFactors: &types.ScalingFactors{
					SearchLevel:       num.DecimalFromFloat(1.2),
					InitialMargin:     num.DecimalFromFloat(1.3),
					CollateralRelease: num.DecimalFromFloat(1.4),
				},
			},
			Instrument: &types.Instrument{
				ID:   "Crypto/ETHUSD/Futures/Dec19",
				Code: "FX:ETHUSD/DEC19",
				Name: "December 2019 ETH vs USD future",
				Metadata: &types.InstrumentMetadata{
					Tags: []string{
						"asset_class:fx/crypto",
						"product:futures",
					},
				},
				Product: &types.InstrumentFuture{
					Future: &types.Future{
						SettlementAsset: "Ethereum/Ether",
						OracleSpecForSettlementPrice: &types.OracleSpec{
							ID:      hex.EncodeToString(crypto.Hash([]byte(ID + "price"))),
							PubKeys: []string{pubKey},
							Filters: []*types.OracleSpecFilter{
								{
									Key: &types.OracleSpecPropertyKey{
										Name: "prices.ETH.value",
										Type: oraclespb.PropertyKey_TYPE_INTEGER,
									},
									Conditions: []*types.OracleSpecCondition{},
								},
							},
						},
						OracleSpecForTradingTermination: &types.OracleSpec{
							ID:      hex.EncodeToString(crypto.Hash([]byte(ID + "tt"))),
							PubKeys: []string{pubKey},
							Filters: []*types.OracleSpecFilter{
								{
									Key: &types.OracleSpecPropertyKey{
										Name: "trading.terminated",
										Type: oraclespb.PropertyKey_TYPE_BOOLEAN,
									},
									Conditions: []*types.OracleSpecCondition{},
								},
							},
						},
						OracleSpecBinding: &types.OracleSpecBindingForFuture{
							SettlementPriceProperty:    "prices.ETH.value",
							TradingTerminationProperty: "trading.terminated",
						},
					},
				},
			},
			RiskModel: &types.TradableInstrumentLogNormalRiskModel{
				LogNormalRiskModel: &types.LogNormalRiskModel{
					RiskAversionParameter: num.DecimalFromFloat(0.01),
					Tau:                   num.DecimalFromFloat(1.0 / 365.25 / 24),
					Params: &types.LogNormalModelParams{
						Mu:    num.DecimalZero(),
						R:     num.DecimalFromFloat(0.016),
						Sigma: num.DecimalFromFloat(0.09),
					},
				},
			},
		},
		State: types.MarketStateActive,
	}
}

func getEngine(t *testing.T, now time.Time) *snapshotTestData {
	t.Helper()
	// ctrl := gomock.NewController(t)
	cfg := execution.NewDefaultConfig()
	log := logging.NewTestLogger()
	broker := stubs.NewBrokerStub()
	timeService := stubs.NewTimeStub()
	timeService.SetTime(now)
	collateralEngine := collateral.New(log, collateral.NewDefaultConfig(), timeService, broker)
	oracleEngine := oracles.NewEngine(log, oracles.NewDefaultConfig(), timeService, broker)

	epochEngine := epochtime.NewService(log, epochtime.NewDefaultConfig(), broker)
	marketActivityTracker := execution.NewMarketActivityTracker(logging.NewTestLogger(), epochEngine)

	ethAsset := types.Asset{
		ID: "Ethereum/Ether",
		Details: &types.AssetDetails{
			Name:   "Ethereum/Ether",
			Symbol: "Ethereum/Ether",
		},
	}
	collateralEngine.EnableAsset(context.Background(), ethAsset)

	eng := execution.NewEngine(
		log,
		cfg,
		timeService,
		collateralEngine,
		oracleEngine,
		broker,
		stubs.NewStateVar(),
		marketActivityTracker,
		stubs.NewAssetStub(),
	)

	statsData := stats.New(log, stats.NewDefaultConfig())
	config := snp.NewDefaultConfig()
	config.Storage = "memory"
	snapshotEngine, _ := snp.New(context.Background(), &paths.DefaultPaths{}, config, log, timeService, statsData.Blockchain)
	snapshotEngine.AddProviders(eng)
	snapshotEngine.ClearAndInitialise()

	return &snapshotTestData{
		engine:         eng,
		oracleEngine:   oracleEngine,
		snapshotEngine: snapshotEngine,
		timeService:    timeService,
	}
}
