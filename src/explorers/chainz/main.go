package chainz

import (
	"model"
	"crawler"
	"fmt"
	"config"
	"encoding/json"
)


var (
	conf      = config.NewConfig(config.DefaultEnvironment)
	txListUrl = "http://chainz.cryptoid.info/%s/api.dws?q=unspent&active=%s&key=%s"
	txInfoUrl = "http://chainz.cryptoid.info/%s/api.dws?q=txinfo&t=%s&key=%s"
)

type Chainz struct {
	Name string
}

func NewChainz() *Chainz {
	return &Chainz{Name: model.CHAINZ_NAME}
}

func (e *Chainz) GetName() string {
	return e.Name
}

func (e *Chainz) CheckForTx(tx *model.Tx) bool {

	txList, err := crawler.Crawl("GET", fmt.Sprintf(txListUrl, tx.Ticker, tx.Sender, conf.GetString("vendors.chainz")), nil, nil)
	if err != nil {
		return false
	}


	unspentTxs := &UnspentTxs{}
	err = json.Unmarshal(txList,unspentTxs)
	//проверяем исходящий кошелек на отправленные транзакции
	for _, chainzTx := range unspentTxs.UnspentOutputs{
		txInfo, err := crawler.Crawl("GET", fmt.Sprintf(txInfoUrl, tx.Ticker, chainzTx.TxHash, conf.GetString("vendors.chainz")), nil, nil)

		fmt.Println(string(txInfo))
		if err != nil {
			return false
		}
		txInfoObject := &TxInfo{}
		err = json.Unmarshal(txInfo, txInfoObject)
		for _, spentOutput := range txInfoObject.Outputs {
			// сравниваем отправленное значение с кошелька с полученным минус 1 процент
			if tx.Amount >= (spentOutput.Amount - spentOutput.Amount / 100 ){
				return true
			}
		}
	}
	return false
}
