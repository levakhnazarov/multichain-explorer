package logger

import (
	"config"
	"encoding/json"
	"fmt"
	syslog "github.com/RackSec/srslog"
	"log"
	"model"
	"strings"
	"time"
)

type Syslog struct {
	Writer *syslog.Writer
}

func NewLogger(p syslog.Priority) *Syslog {
	confObject := config.NewConfig(config.DefaultEnvironment)

	hostname := strings.Split(confObject.GetString("ecom.app.log_server"), ":")

	if len(hostname) != 2 {
		log.Printf("failed to split host string, required format = hostname:port")
	}

	//ip, err := net.ResolveIPAddr("", hostname[0])
	//if err != nil {
	//	log.Printf("failed to resolve ip, err %s", err.Error())
	//}
	//remoteTcpString := ip.IP.String() + ":" + hostname[1]
	remoteTcpString := ""

	w, err := syslog.Dial("tcp", remoteTcpString, p, confObject.GetString("ecom.app.name"))
	if err != nil {
		//log.Printf("failed to dial syslog, err: %s", err.Error())
	}

	return &Syslog{
		Writer: w,
	}
}

func (s *Syslog) Log(message []string) {

	if len(message) == 3 {
		message := &model.LogMessage{
			Cmd: message[0],
			Msg: message[1],
			Err: message[2],
			Tm:  time.Now(),
		}
		serializedMessage, err := json.Marshal(message)
		if err == nil {
			fmt.Println(string(serializedMessage))
			//s.Writer.Alert(string(serializedMessage))

		}

	} else {
		log.Fatalf("not enough args to call to syslog, expected: %d", 3)
	}
}
