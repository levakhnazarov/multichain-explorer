package chainz

type UnspentTxs struct {
	UnspentOutputs []struct {
		TxHash        string `json:"tx_hash"`
		TxOuputN      int    `json:"tx_ouput_n"`
		Value         string `json:"value"`
		Confirmations int    `json:"confirmations"`
		Script        string `json:"script"`
	} `json:"unspent_outputs"`
}

type TxInfo struct {
	Hash          string  `json:"hash"`
	Block         int     `json:"block"`
	Index         int     `json:"index"`
	Timestamp     int     `json:"timestamp"`
	Confirmations int     `json:"confirmations"`
	Fees          float64 `json:"fees"`
	TotalInput    float64 `json:"total_input"`
	Inputs        []struct {
		Addr         string  `json:"addr"`
		Amount       float64 `json:"amount"`
		ReceivedFrom struct {
			Tx string `json:"tx"`
			N  int    `json:"n"`
		} `json:"received_from"`
	} `json:"inputs"`
	TotalOutput float64 `json:"total_output"`
	Outputs     []struct {
		Addr   string  `json:"addr"`
		Amount float64 `json:"amount"`
		Script string  `json:"script"`
	} `json:"outputs"`
}