package explorers

import (
	"explorers/chainz"
	"explorers/etherscan"
	"fmt"

	"model"
)

type AbstractExplorer interface {
	GetName() string
	CheckForTx(tx *model.Tx) bool
}

type Aggregator struct {
	Explorers []interface{}
	Name      string
}

func NewAggregator() *Aggregator {
	aggr := &Aggregator{}
	aggr.Explorers = append(aggr.Explorers, chainz.NewChainz())
	aggr.Explorers = append(aggr.Explorers, etherscan.NewEtherScan())
	return aggr
}

func (a *Aggregator) CheckIfTxExist(tx *model.Tx) error {
	for _, expl := range a.Explorers {
		c, ok := expl.(AbstractExplorer)

		if !ok {
			return fmt.Errorf("can't receive interface type")
		}
		//fmt.Println(string(txList))
		if c.GetName() == tx.Explorer {
			if c.CheckForTx(tx) {
				return nil
			} else {
				return fmt.Errorf("can't process tx")
			}
		}

	}

	return nil

}
