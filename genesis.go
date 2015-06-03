package core

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

/*
 * This is the special genesis block.
 */

var ZeroHash256 = make([]byte, 32)
var ZeroHash160 = make([]byte, 20)
var ZeroHash512 = make([]byte, 64)

func GenesisBlock(db common.Database) *types.Block {
	genesis := types.NewBlock(common.Hash{}, common.Address{}, common.Hash{}, params.GenesisDifficulty, 42, nil)
	genesis.Header().Number = common.Big0
	genesis.Header().GasLimit = params.GenesisGasLimit
	genesis.Header().GasUsed = common.Big0
	genesis.Header().Time = 0

	genesis.Td = common.Big0

	genesis.SetUncles([]*types.Header{})
	genesis.SetTransactions(types.Transactions{})
	genesis.SetReceipts(types.Receipts{})

	var accounts map[string]struct {
		Balance string
		Code    string
	}
	err := json.Unmarshal(GenesisData, &accounts)
	if err != nil {
		fmt.Println("enable to decode genesis json data:", err)
		os.Exit(1)
	}

	statedb := state.New(genesis.Root(), db)
	for addr, account := range accounts {
		codedAddr := common.Hex2Bytes(addr)
		accountState := statedb.CreateAccount(common.BytesToAddress(codedAddr))
		accountState.SetBalance(common.Big(account.Balance))
		accountState.SetCode(common.FromHex(account.Code))
		statedb.UpdateStateObject(accountState)
	}
	statedb.Sync()
	genesis.Header().Root = statedb.Root()
	genesis.Td = params.GenesisDifficulty

	return genesis
}

var GenesisData = []byte(`{
  "aced1ce9bb193d4270acf8738942ac7d008f22b4":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"},
  "f1abb0d8af6f3de43cf05eb3d9458c95e79f30a0":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "ba5e10071204f37931769d7afa454bc82e1eb4cd":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "a1e5acc2a0c6efa671d9e27c4faf60f22fc50de0":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "bade1f4d04f56f13381b6e3dc0e78479fe2563c2":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "e1fd1e0689bc35a4b7e531f50f96d77b02149dbc":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "dad1005e6487a99422b22a9db665e01148add52b":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "fa15efeef38db9ff6f7eb862e3402c66ee0e7951":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "fadedcabf08aead902061c146c5458e5bea1ce9f":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "5eedfab1e5e0ed81674dc6b503c6c89d413ccbc9":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}, 
  "1dea5c01dd17dee05b08c89e0753722f52f6d2f1":
      {"balance": "1606938044258990275541962092341162602522202993782792835301376"}
  }`)
