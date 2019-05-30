package tailf

import (
	"github.com/hpcloud/tail"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

type TextMsg struct {
	Msg string
	Topic string 
}

type TailObjMgr struct {
	tailsObjs []*TailObj
	msgChan chan *TextMsg
}

var (
	tailObjMgr *TailObjMgr
)

func GetOneLine()(msg *TextMsg) {
	msg = <- tailObjMgr.msgChan
	return
}

func InitTail(conf []CollectConf,chanSize int) (err error) {
	if len(conf) == 0 {
		err = fmt.Errorf("invalid config for log collect",conf)
		return 
	}
	tailObjMgr = &TailObjMgr{
		msgChan : make(chan *TextMsg,chanSize),
	}
	for _,v := range conf {
		obj := &TailObj{
			conf : v,
		}
		tails,errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen    :   true,
			Follow    :   true,
			MustExist :   false,
			Poll      :   true ,
		}) 
		if errTail != nil {
			err = errTail
			return 
		}
		obj.tail = tails
		tailObjMgr.tailsObjs = append(tailObjMgr.tailsObjs, obj)

		go readFromTail(obj)
	}
	return 
}

func readFromTail(tailObj *TailObj) {
	for true {
		line,ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen,filename:%s\n",tailObj.tail.Filename)
			time.Sleep(100*time.Millisecond)
			continue
		}

		textMsg := &TextMsg {
			Msg : line.Text,
			Topic : tailObj.conf.Topic,
		}
		tailObjMgr.msgChan<- textMsg
	}
}