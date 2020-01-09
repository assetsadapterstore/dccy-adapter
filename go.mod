module github.com/assetsadapterstore/dccy-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/eosio-adapter v1.3.0
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet v1.7.0
	github.com/eoscanada/eos-go v0.8.16
	github.com/imroc/req v0.2.4
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/tidwall/gjson v1.3.5
)

replace github.com/eoscanada/eos-go => github.com/blocktree/eos-go v0.8.13-blocktree

//replace github.com/blocktree/eosio-adapter => ../eosio-adapter
