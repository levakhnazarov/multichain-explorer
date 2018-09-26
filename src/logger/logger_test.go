package logger

import (
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {

	hostname := strings.Split("www.google.com:80", ":")

	if len(hostname) != 2 {
		log.Fatalf("failed to split host string, required format = hostname:port")
	}

	ip, err := net.ResolveIPAddr("", "www.golang.org")
	if err != nil {
		log.Fatalf("failed to resolve ip, err %s", err.Error())
	}

	fmt.Print(ip.IP.String())

}
