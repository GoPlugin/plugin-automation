package ocr_test

import (
	"context"
	"io"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/goplugin/plugin-automation/pkg/v3/types"
	"github.com/goplugin/plugin-automation/tools/simulator/config"
	"github.com/goplugin/plugin-automation/tools/simulator/simulate/chain"
	"github.com/goplugin/plugin-automation/tools/simulator/simulate/ocr"
	"github.com/goplugin/plugin-automation/tools/simulator/util"
	ocr2keepers "github.com/goplugin/plugin-common/pkg/types/automation"
)

func TestReportTracker(t *testing.T) {
	t.Parallel()

	logger := log.New(io.Discard, "", 0)
	conf := config.Blocks{
		Genesis:  new(big.Int).SetInt64(1),
		Cadence:  config.Duration(100 * time.Millisecond),
		Jitter:   config.Duration(0),
		Duration: 10,
	}

	upkeepID := util.NewUpkeepID(big.NewInt(8).Bytes(), uint8(types.ConditionTrigger))
	workID := util.UpkeepWorkID(
		upkeepID,
		ocr2keepers.NewLogTrigger(
			ocr2keepers.BlockNumber(5),
			[32]byte{},
			nil))

	report1, err := util.EncodeCheckResultsToReportBytes([]ocr2keepers.CheckResult{
		{
			UpkeepID: upkeepID,
			WorkID:   workID,
		},
	})

	require.NoError(t, err)

	transmits := []chain.TransmitEvent{
		{
			SendingAddress: "test",
			BlockNumber:    big.NewInt(1),
			Report:         report1,
		},
	}

	broadcaster := chain.NewBlockBroadcaster(conf, 1, logger, nil, loadTransmitsAt(transmits, 2))
	listener := chain.NewListener(broadcaster, logger)

	tracker := ocr.NewReportTracker(listener, logger)

	<-broadcaster.Start()
	broadcaster.Stop()

	evts, err := tracker.GetLatestEvents(context.Background())

	require.NoError(t, err)
	assert.Len(t, evts, 1)
}

func loadTransmitsAt(transmits []chain.TransmitEvent, atBlock int64) func(*chain.Block) {
	return func(block *chain.Block) {
		if block.Number.Cmp(new(big.Int).SetInt64(atBlock)) == 0 {
			block.Transactions = append(block.Transactions, chain.PerformUpkeepTransaction{
				Transmits: transmits,
			})
		}
	}
}
