package main

import(
	"fmt"
	"github.com/astaxie/beego/config"
	"errors"
	"logagent/tailf"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath  string
	chanSize int
	conllectConf []tailf.CollectConf
	kafkaAddr string
}



func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collent::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collent::topic")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::Topic")
		return
	}

	appConfig.conllectConf = append(appConfig.conllectConf,cc)
	return 
}

func loadConf(confType,filename string) (err error) {
	conf,err := config.NewConfig(confType,filename)
	if err != nil {
		fmt.Println("new config failed,err:",err)
		return 
	}

	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "D:/goproject/logs/logagent.log"
	}

	appConfig.chanSize,err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}
	
	appConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed,err:%v\n", err)
		return
	}
	return 
}