package rpc

import (
	"fmt"
	"encoding/json"
	"github.com/Leondroids/go-ethereum-rpc/types"
)

type SyncStatus struct {
	StartingBlock int64 `json:"startingBlock"`
	HighestBlock  int64 `json:"highestBlock"`
	CurrentBlock  int64 `json:"currentBlock"`
}

type SyncStatusRaw struct {
	StartingBlock string `json:"startingBlock"`
	HighestBlock  string `json:"highestBlock"`
	CurrentBlock  string `json:"currentBlock"`
}

func (client *Client) RequestSyncStatus(method string, params ...interface{}) (*SyncStatus, error) {
	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no transaction found for %v", params)
	}

	js, err := json.Marshal(response.Result)

	if err != nil {
		return nil, err
	}

	return new(SyncStatusRaw).fromJSON(js)
}

func (ssr *SyncStatusRaw) fromJSON(js []byte) (*SyncStatus, error) {
	err := json.Unmarshal(js, ssr)
	if err != nil {
		return nil, err
	}
	return ssr.ToSyncStatus()
}

func (ssr *SyncStatusRaw) ToSyncStatus() (*SyncStatus, error) {

	startingBlock, err := types.NewHexString(ssr.StartingBlock)
	if err != nil {
		return nil, err
	}
	highestBlock, err := types.NewHexString(ssr.HighestBlock)
	if err != nil {
		return nil, err
	}
	currentBlock, err := types.NewHexString(ssr.CurrentBlock)
	if err != nil {
		return nil, err
	}

	return &SyncStatus{
		StartingBlock: startingBlock.Int64(),
		HighestBlock:  highestBlock.Int64(),
		CurrentBlock:  currentBlock.Int64(),
	}, nil
}
