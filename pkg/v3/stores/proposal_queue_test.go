package stores

import (
	"testing"

	"github.com/goplugin/plugin-automation/pkg/v3/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ocr2keepers "github.com/goplugin/plugin-common/pkg/types/automation"
)

func TestProposalQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name      string
		initials  []ocr2keepers.CoordinatedBlockProposal
		toEnqueue []ocr2keepers.CoordinatedBlockProposal
		size      int
	}{
		{
			"add to empty queue",
			[]ocr2keepers.CoordinatedBlockProposal{},
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.LogTrigger, []byte{0x1}),
					WorkID:   "0x1",
				},
			},
			1,
		},
		{
			"add to non-empty queue",
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.LogTrigger, []byte{0x1}),
					WorkID:   "0x1",
				},
			},
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.LogTrigger, []byte{0x2}),
					WorkID:   "0x2",
				},
			},
			2,
		},
		{
			"add existing",
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.LogTrigger, []byte{0x1}),
					WorkID:   "0x1",
				},
			},
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.LogTrigger, []byte{0x1}),
					WorkID:   "0x1",
				},
			},
			1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			q := NewProposalQueue(func(uid ocr2keepers.UpkeepIdentifier) types.UpkeepType {
				return types.UpkeepType(uid[15])
			})

			require.NoError(t, q.Enqueue(tc.initials...))
			require.NoError(t, q.Enqueue(tc.toEnqueue...))
			require.Equal(t, tc.size, q.Size())
		})
	}
}

func TestProposalQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name         string
		toEnqueue    []ocr2keepers.CoordinatedBlockProposal
		dequeueType  types.UpkeepType
		dequeueCount int
		expected     []ocr2keepers.CoordinatedBlockProposal
	}{
		{
			"empty queue",
			[]ocr2keepers.CoordinatedBlockProposal{},
			types.LogTrigger,
			1,
			[]ocr2keepers.CoordinatedBlockProposal{},
		},
		{
			"happy path log trigger",
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.LogTrigger, []byte{0x1}),
					WorkID:   "0x1",
				},
				{
					UpkeepID: upkeepId(types.ConditionTrigger, []byte{0x1}),
					WorkID:   "0x2",
				},
			},
			types.LogTrigger,
			2,
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.LogTrigger, []byte{0x1}),
					WorkID:   "0x1",
				},
			},
		},
		{
			"happy path log trigger",
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.LogTrigger, []byte{0x1}),
					WorkID:   "0x1",
				},
				{
					UpkeepID: upkeepId(types.ConditionTrigger, []byte{0x1}),
					WorkID:   "0x2",
				},
			},
			types.ConditionTrigger,
			2,
			[]ocr2keepers.CoordinatedBlockProposal{
				{
					UpkeepID: upkeepId(types.ConditionTrigger, []byte{0x1}),
					WorkID:   "0x2",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			q := NewProposalQueue(func(uid ocr2keepers.UpkeepIdentifier) types.UpkeepType {
				return types.UpkeepType(uid[15])
			})
			for _, p := range tc.toEnqueue {
				err := q.Enqueue(p)
				assert.NoError(t, err)
			}
			results, err := q.Dequeue(tc.dequeueType, tc.dequeueCount)
			require.NoError(t, err)
			require.Equal(t, len(tc.expected), len(results))

			for i := range tc.expected {
				require.Equal(t, tc.expected[i].WorkID, results[i].WorkID)
			}
		})
	}
}

func upkeepId(utype types.UpkeepType, rand []byte) ocr2keepers.UpkeepIdentifier {
	id := [32]byte{}
	id[15] = byte(utype)
	copy(id[16:], rand)
	return ocr2keepers.UpkeepIdentifier(id)
}
