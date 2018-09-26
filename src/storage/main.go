package storage

import (
	"config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"model"
)

//TODO singleton
var conf = config.NewConfig(config.DefaultEnvironment)

var (
	dbinfo          = fmt.Sprintf("%s:%s@%s/%s", conf.GetString("mysql.user"), conf.GetString("mysql.password"), conf.GetString("mysql.host"), conf.GetString("mysql.database"))
	getCoinByTicker = `SELECT explorer, name FROM Currencies where ticker = ?;`
)

func GetSingleCoinExplorer(ticker string) (*model.CoinFromDb, error) {
	fmt.Println(ticker)
	respCoin := &model.CoinFromDb{}
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmtOut, err := db.Prepare(getCoinByTicker)
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	var explorer, name string

	err = stmtOut.QueryRow(ticker).Scan(&explorer, &name)
	if err != nil {
		return nil, err
	}
	respCoin.Explorer = explorer
	respCoin.Name = name
	respCoin.Ticker = ticker

	return respCoin, nil

}
