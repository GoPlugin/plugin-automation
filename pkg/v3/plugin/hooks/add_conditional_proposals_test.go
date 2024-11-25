package hooks

import (
	"bytes"
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	ocr2keepers "github.com/goplugin/plugin-automation/pkg/v3"
	"github.com/goplugin/plugin-automation/pkg/v3/types"
	common "github.com/goplugin/plugin-common/pkg/types/automation"
)

func TestAddConditionalSamplesHook_RunHook(t *testing.T) {
	for _, tc := range []struct {
		name             string
		metadata         types.MetadataStore
		coordinator      types.Coordinator
		proposals        []common.CoordinatedBlockProposal
		limit            int
		src              [16]byte
		wantNumProposals int
		expectErr        bool
		wantErr          error
	}{
		{
			name: "proposals aren't filtered and are added to the observation",
			metadata: &mockMetadataStore{
				ViewConditionalProposalFn: func() []common.CoordinatedBlockProposal {
					return []common.CoordinatedBlockProposal{
						{
							WorkID: "workID1",
						},
					}
				},
			},
			coordinator: &mockCoordinator{
				FilterProposalsFn: func(proposals []common.CoordinatedBlockProposal) ([]common.CoordinatedBlockProposal, error) {
					assert.Equal(t, 1, len(proposals))
					return proposals, nil
				},
			},
			limit:            5,
			src:              [16]byte{1},
			wantNumProposals: 1,
		},
		{
			name: "proposals are filtered and are added to the observation",
			metadata: &mockMetadataStore{
				ViewConditionalProposalFn: func() []common.CoordinatedBlockProposal {
					return []common.CoordinatedBlockProposal{
						{
							WorkID: "workID1",
						},
						{
							WorkID: "workID2",
						},
					}
				},
			},
			coordinator: &mockCoordinator{
				FilterProposalsFn: func(proposals []common.CoordinatedBlockProposal) ([]common.CoordinatedBlockProposal, error) {
					assert.Equal(t, 2, len(proposals))
					return proposals[:1], nil
				},
			},
			limit:            5,
			src:              [16]byte{1},
			wantNumProposals: 1,
		},
		{
			name: "proposals are appended to the existing proposals in observation",
			metadata: &mockMetadataStore{
				ViewConditionalProposalFn: func() []common.CoordinatedBlockProposal {
					return []common.CoordinatedBlockProposal{
						{
							WorkID: "workID1",
						},
					}
				},
			},
			coordinator: &mockCoordinator{
				FilterProposalsFn: func(proposals []common.CoordinatedBlockProposal) ([]common.CoordinatedBlockProposal, error) {
					assert.Equal(t, 1, len(proposals))
					return proposals, nil
				},
			},
			proposals:        []common.CoordinatedBlockProposal{{WorkID: "workID2"}},
			limit:            5,
			src:              [16]byte{1},
			wantNumProposals: 2,
		},
		{
			name: "proposals aren't filtered but are limited and are added to the observation",
			metadata: &mockMetadataStore{
				ViewConditionalProposalFn: func() []common.CoordinatedBlockProposal {
					return []common.CoordinatedBlockProposal{
						{
							WorkID: "workID1",
						},
						{
							WorkID: "workID2",
						},
						{
							WorkID: "workID3",
						},
						{
							WorkID: "workID4",
						},
					}
				},
			},
			coordinator: &mockCoordinator{
				FilterProposalsFn: func(proposals []common.CoordinatedBlockProposal) ([]common.CoordinatedBlockProposal, error) {
					assert.Equal(t, 4, len(proposals))
					return proposals, nil
				},
			},
			limit:            2,
			src:              [16]byte{0},
			wantNumProposals: 2,
		},
		{
			name: "if an error is encountered filtering proposals, an error is returned",
			metadata: &mockMetadataStore{
				ViewConditionalProposalFn: func() []common.CoordinatedBlockProposal {
					return []common.CoordinatedBlockProposal{
						{
							WorkID: "workID1",
						},
						{
							WorkID: "workID2",
						},
						{
							WorkID: "workID3",
						},
						{
							WorkID: "workID4",
						},
					}
				},
			},
			coordinator: &mockCoordinator{
				FilterProposalsFn: func(proposals []common.CoordinatedBlockProposal) ([]common.CoordinatedBlockProposal, error) {
					return nil, errors.New("filter proposals boom")
				},
			},
			limit:            2,
			src:              [16]byte{0},
			wantNumProposals: 0,
			expectErr:        true,
			wantErr:          errors.New("filter proposals boom"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var logBuf bytes.Buffer
			logger := log.New(&logBuf, "", 0)
			processor := NewAddConditionalProposalsHook(tc.metadata, tc.coordinator, logger)
			observation := &ocr2keepers.AutomationObservation{
				UpkeepProposals: tc.proposals,
			}
			err := processor.RunHook(observation, tc.limit, tc.src)
			if tc.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.wantNumProposals, len(observation.UpkeepProposals))
		})
	}
}
