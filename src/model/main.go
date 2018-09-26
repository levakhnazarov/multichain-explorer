package model

import "time"

type LogMessage struct {
	Cmd string
	Msg string
	Err string
	Tm  time.Time
}

type JsonResponse struct {
	Success bool   `json:"success"`
	Comment string `json:"comment,omitempty"`
}

type CoinFromDb struct {
	Explorer string `json:"explorer,omitempty"`
	Name     string `json:"name,omitempty"`
	Ticker   string `json:"ticker,omitempty"`
}

type Tx struct {
	Amount                                        float64
	Timestamp, Explorer, Sender, Receiver, Ticker string
}

type CrawlParam struct {
	Key   string
	Value string
}
