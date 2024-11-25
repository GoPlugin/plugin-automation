package postprocessors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goplugin/plugin-automation/pkg/v3/stores"
	"github.com/goplugin/plugin-automation/pkg/v3/types"
	ocr2keepers "github.com/goplugin/plugin-common/pkg/types/automation"
)

func TestMetadataAddSamples(t *testing.T) {
	ch := make(chan ocr2keepers.BlockHistory)
	ms, err := stores.NewMetadataStore(&mockBlockSubscriber{ch: ch}, func(uid ocr2keepers.UpkeepIdentifier) types.UpkeepType {
		return types.ConditionTrigger
	})
	assert.NoError(t, err)

	values := []ocr2keepers.CheckResult{
		{
			Eligible: true,
			UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{1}),
			WorkID:   "workID1",
		},
		{
			Eligible: true,
			UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{2}),
			WorkID:   "workID2",
		},
		{
			Eligible: false,
			UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{3}),
			WorkID:   "workID3",
		},
	}

	pp := NewAddProposalToMetadataStorePostprocessor(ms)
	err = pp.PostProcess(context.Background(), values, []ocr2keepers.UpkeepPayload{
		{
			UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{1}),
			WorkID:   "workID1",
		},
		{
			UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{2}),
			WorkID:   "workID2",
		},
		{
			UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{3}),
			WorkID:   "workID3",
		},
	})

	assert.NoError(t, err, "no error expected from post processor")

	assert.Equal(t, 2, len(ms.ViewProposals(types.ConditionTrigger)))
}

type mockBlockSubscriber struct {
	ch      chan ocr2keepers.BlockHistory
	StartFn func(ctx context.Context) error
	CloseFn func() error
}

func (_m *mockBlockSubscriber) Subscribe() (int, chan ocr2keepers.BlockHistory, error) {
	return 0, _m.ch, nil
}

func (_m *mockBlockSubscriber) Unsubscribe(int) error {
	return nil
}

func (_m *mockBlockSubscriber) Start(ctx context.Context) error {
	return _m.StartFn(ctx)
}

func (_m *mockBlockSubscriber) Close() error {
	return _m.CloseFn()
}
