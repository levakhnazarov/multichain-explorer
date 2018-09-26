package crawler

import (
	"bytes"
	"io/ioutil"
	"model"
	"net/http"

	"fmt"
	"net/http/httputil"
)

func Crawl(method, urlString string, body []byte, headers []model.CrawlParam) (respBody []byte, err error) {

	req, err := http.NewRequest(method, urlString, bytes.NewBuffer(body))

	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	log, err := httputil.DumpRequest(req, true)
	fmt.Print(string(log))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	responce, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return responce, nil

}
