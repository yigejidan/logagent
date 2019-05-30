package main

import (
	"github.com/astaxie/beego/logs"
	"logagent/tailf"
	"time"
	"logagent/kafka"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err := sendToKafKa(msg)
		if err != nil {
			logs.Error("send to kafka failed,err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return 
}

func sendToKafKa(msg *tailf.TextMsg)(err error) {
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}