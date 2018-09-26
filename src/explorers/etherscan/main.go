package etherscan

import (
	"config"
	"crawler"
	"fmt"
	"model"
	"encoding/json"
	"os"
)

type Etherscan struct {
	Name string
}

var (
	conf      = config.NewConfig(config.DefaultEnvironment)
	txListUrl = "http://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=%d&endblock=99999999&sort=asc&apikey=%s"
)

func NewEtherScan() *Etherscan {

	return &Etherscan{ Name: model.ETHERSCAN_NAME }
}

func (e *Etherscan) GetName() string {
	return e.Name
}
func (e *Etherscan) CheckForTx(tx *model.Tx) bool {

	txList, err := crawler.Crawl("GET", fmt.Sprintf(txListUrl, tx.Receiver, 5008930, conf.GetString("vendors.etherscan")), nil, nil)
	if err != nil {
		return false
	}

	txListObject := &TxList{}
	err = json.Unmarshal(txList, txListObject)
	if err != nil {
		return false
	}
	f, err := os.Create("response.json")
	if err != nil {
		return false
	}
	defer f.Close()
	n2, err := f.Write(txList)

	fmt.Println(n2)
	return true
}
