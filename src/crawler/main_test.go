package crawler

import (
	"net/http"
	"testing"
)

func TestCrawler_Crawl(t *testing.T) {

	request, err := http.NewRequest("GET", "http://localhost", nil)

	if err == nil {
		err, data := Crawl(request)
		if err != nil {
			t.Log(err)
		} else {
			t.Log(data.StatusCode)
		}

	}

}
