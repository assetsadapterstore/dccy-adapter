/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package dccy

import (
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/tidwall/gjson"
)

type WalletManager struct {
	openwallet.AssetsAdapterBase

	Api             *eos.API                        // 节点客户端
	Config          *WalletConfig                   // 节点配置
	Decoder         openwallet.AddressDecoder       //地址编码器
	TxDecoder       openwallet.TransactionDecoder   //交易单编码器
	Log             *log.OWLogger                   //日志工具
	ContractDecoder openwallet.SmartContractDecoder //智能合约解析器
	Blockscanner    openwallet.BlockScanner         //区块扫描器
	CacheManager    openwallet.ICacheManager        //缓存管理器
	client          *Client                         //RPC客户端
}

func NewWalletManager(cacheManager openwallet.ICacheManager) *WalletManager {
	wm := WalletManager{}
	wm.Config = NewConfig(Symbol)
	wm.Blockscanner = NewDCCYBlockScanner(&wm)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.TxDecoder = NewTransactionDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	wm.ContractDecoder = NewContractDecoder(&wm)
	wm.CacheManager = cacheManager

	ecc.PublicKeyPrefixs = []string{"EOS", "DCCY"}

	return &wm
}

func (wm *WalletManager) GetCurrencyBalance(account eos.AccountName, symbol string, code eos.AccountName, publisher string) ([]eos.Asset, error) {
	//curl -X POST --url http://IP/v1/chain/get_table_rows -d '{"code":"eosio.token","json":true,"table":"accounts","scope":"accxxx"}'
	//{"rows":[{"primary":0,"balance":{"quantity":"100.0000 DCCY","contract":"eosio"}}],"more":false}
	param := eos.GetTableRowsRequest{
		Code:       string(code),
		Scope:      string(account),
		Table:      "accounts",
		JSON:       true,
	}
	out, err := wm.Api.GetTableRows(param)
	if err != nil {
		return nil, err
	}
	rows := gjson.ParseBytes(out.Rows)
	assets := make([]eos.Asset, 0)
	if rows.IsArray() {
		for _, obj := range rows.Array()  {
			balance := obj.Get("balance.quantity").String()
			contract := obj.Get("balance.contract").String()
			a, e := eos.NewAsset(balance)
			if e != nil {
				return nil, e
			}
			if a.Symbol.Symbol == symbol && contract == publisher {
				assets = append(assets, a)
			}
		}
	}

	return assets, nil
}