// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package native

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MarshalJSON marshals as JSON.
func (c callFrame) MarshalJSON() ([]byte, error) {
	type callFrame0 struct {
		BeforeEVMTransfers *[]arbitrumTransfer `json:"beforeEVMTransfers,omitempty"`
		AfterEVMTransfers  *[]arbitrumTransfer `json:"afterEVMTransfers,omitempty"`
		Type               vm.OpCode           `json:"-"`
		From               common.Address      `json:"from"`
		Gas                uint64              `json:"gas"`
		GasUsed            uint64              `json:"gasUsed"`
		To                 common.Address      `json:"to,omitempty" rlp:"optional"`
		Input              []byte              `json:"input" rlp:"optional"`
		Output             []byte              `json:"output,omitempty" rlp:"optional"`
		Error              string              `json:"error,omitempty" rlp:"optional"`
		RevertReason       string              `json:"revertReason,omitempty"`
		Calls              []callFrame         `json:"calls,omitempty" rlp:"optional"`
		Logs               []callLog           `json:"logs,omitempty" rlp:"optional"`
		Value              *big.Int            `json:"value,omitempty" rlp:"optional"`
		TypeString	   string              `json:"type"`
	}
	var enc callFrame0
	enc.BeforeEVMTransfers = c.BeforeEVMTransfers
	enc.AfterEVMTransfers = c.AfterEVMTransfers
	enc.Type = c.Type
	enc.From = c.From
	enc.Gas = c.Gas
	enc.GasUsed = c.GasUsed
	enc.To = c.To
	enc.Input = c.Input
	enc.Output = c.Output
	enc.Error = c.Error
	enc.RevertReason = c.RevertReason
	enc.Calls = c.Calls
	enc.Logs = c.Logs
	enc.Value = c.Value
	enc.TypeString = c.TypeString()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (c *callFrame) UnmarshalJSON(input []byte) error {
	type callFrame0 struct {
		BeforeEVMTransfers *[]arbitrumTransfer `json:"beforeEVMTransfers,omitempty"`
		AfterEVMTransfers  *[]arbitrumTransfer `json:"afterEVMTransfers,omitempty"`
		Type               *vm.OpCode          `json:"-"`
		From               *common.Address     `json:"from"`
		Gas                *uint64             `json:"gas"`
		GasUsed            *uint64             `json:"gasUsed"`
		To                 *common.Address     `json:"to,omitempty" rlp:"optional"`
		Input              []byte              `json:"input" rlp:"optional"`
		Output             []byte              `json:"output,omitempty" rlp:"optional"`
		Error              *string             `json:"error,omitempty" rlp:"optional"`
		RevertReason       *string             `json:"revertReason,omitempty"`
		Calls              []callFrame         `json:"calls,omitempty" rlp:"optional"`
		Logs               []callLog           `json:"logs,omitempty" rlp:"optional"`
		Value              *big.Int            `json:"value,omitempty" rlp:"optional"`
	}
	var dec callFrame0
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BeforeEVMTransfers != nil {
		c.BeforeEVMTransfers = dec.BeforeEVMTransfers
	}
	if dec.AfterEVMTransfers != nil {
		c.AfterEVMTransfers = dec.AfterEVMTransfers
	}
	if dec.Type != nil {
		c.Type = *dec.Type
	}
	if dec.From != nil {
		c.From = *dec.From
	}
	if dec.Gas != nil {
		c.Gas = *dec.Gas
	}
	if dec.GasUsed != nil {
		c.GasUsed = *dec.GasUsed
	}
	if dec.To != nil {
		c.To = *dec.To
	}
	if dec.Input != nil {
		c.Input = dec.Input
	}
	if dec.Output != nil {
		c.Output = dec.Output
	}
	if dec.Error != nil {
		c.Error = *dec.Error
	}
	if dec.RevertReason != nil {
		c.RevertReason = *dec.RevertReason
	}
	if dec.Calls != nil {
		c.Calls = dec.Calls
	}
	if dec.Logs != nil {
		c.Logs = dec.Logs
	}
	if dec.Value != nil {
		c.Value = dec.Value
	}
	return nil
}
