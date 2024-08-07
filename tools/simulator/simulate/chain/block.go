package chain

import (
	"math/big"

	"github.com/goplugin/plugin-libocr/offchainreporting2plus/types"
)

type Block struct {
	Hash         [32]byte
	Number       *big.Int
	Transactions []interface{}
}

type Log struct {
	TxHash       [32]byte
	BlockNumber  *big.Int
	BlockHash    [32]byte
	Idx          uint32
	TriggerValue string
}

type OCR3ConfigTransaction struct {
	Config types.ContractConfig
}

type PerformUpkeepTransaction struct {
	Transmits []TransmitEvent
}

type UpkeepCreatedTransaction struct {
	Upkeep SimulatedUpkeep
}

// below this line should not be in this package
type UpkeepType int

const (
	ConditionalType UpkeepType = iota
	LogTriggerType
)

type SimulatedUpkeep struct {
	ID             *big.Int
	CreateInBlock  *big.Int
	UpkeepID       [32]byte
	Type           UpkeepType
	AlwaysEligible bool
	EligibleAt     []*big.Int
	TriggeredBy    string
	CheckData      []byte
	Expected       bool
}

type SimulatedLog struct {
	TriggerAt    *big.Int
	TriggerValue string
}

type TransmitEvent struct {
	SendingAddress string
	Report         []byte
	Hash           [32]byte
	Round          uint64
	BlockNumber    *big.Int
	BlockHash      [32]byte
}
