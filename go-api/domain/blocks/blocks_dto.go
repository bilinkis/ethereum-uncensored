package blocks

import (
	"block-tracker/utils/rest_errors"
	"strings"
)

type Block struct {
	Number string `json:"number"`
	Miner string `json:"miner"`
	Orphan bool `json:"orphan"`
}

func (block *Block) Validate() rest_errors.RestErr {
	if block.Number == "" || block.Miner == "" {
		return rest_errors.NewBadRequestError("invalid block data")
	}

	block.Number = TrStr(block.Number)
	block.Miner = TrStr(block.Miner)

	return nil
}

func TrStr(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

type InfuraHeaderResponse struct {
	ExcOpt bool `json:"execution_optimistic"`
	Data []IHRData `json:"data"`
}

type IHRData struct {
	Root string `json:"execution_optimistic"`
	Canonical bool `json:"canonical"`
	Header IHRDHeader `json:"header"`
}

type IHRDHeader struct {
	Message IHRDHMessage `json:"message"`
	Signature string `json:"signature"`
}

type IHRDHMessage struct {
	Slot string `json:"slot"`
	ProposerIndex string `json:"proposer_index"`
	ParentRoot string `json:"parent_root"`
	StateRoot string `json:"state_root"`
	BodyRoot string `json:"body_root"`
}