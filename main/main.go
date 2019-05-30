package main

import(
	"fmt"
	"github.com/astaxie/beego/logs"
	"logagent/tailf"
	"logagent/kafka"
)

func main(){
	filename := "D:/goproject/conf/logagent.conf"
	err := loadConf("ini",filename)
	if err != nil {
		fmt.Printf("load conf failed,err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed,err:%v\n", err)
		panic("load logger failed")
		return 
	}

	
	logs.Debug("load conf succ,config:%v", appConfig)

	err = tailf.InitTail(appConfig.conllectConf,appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed,err:%v", err)
		return 
	}
	logs.Debug("initialize tailf succ")

	err = kafka.InitKafKa(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init tail failed,err:%v", err)
		return 
	}

	logs.Debug("initialize all succ")
	// go func() {
	// 	var count int
	// 	for {
	// 		count++
	// 		logs.Debug("test for logger %d", count)
	// 		time.Sleep(time.Millisecond*1000)
	// 	}
	// }()

	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed ,err:%v", err)
		return 
	}

	logs.Info("program exited")
}