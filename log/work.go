package log

import (
	"Cave-CTF/config"
	"fmt"
	"time"
)

var piping chan info

func init() {
	config.Config.Read()
	day = time.Now().YearDay()
	logFileName = "cave.log"
	piping = make(chan info)
	openLogFile()
	go work()
}

// 处理信息
func work() {
	data := <-piping
	var level string
	switch data.Level {
	case DEBUG:
		level = "DEBUG"
	case INFO:
		level = "INFO"
	case WARNING:
		level = "WARNING"
	case ERROR:
		level = "ERROR"
	case CRITICAL:
		level = "CRITICAL"
	}
	str := fmt.Sprintf("[Cave-CTF] %s %s Code:%d Msg:%s Error:%s\n",
		data.Date.Format("2006-01-02 15:04:05"),
		level,
		data.Code,
		data.Msg,
		data.Err,
	)
	outDisplay(str)
	outFile(str)
}

//func work_0() {
//	data := <-piping
//}
//
//func work_1() {
//	data := <-piping
//}
//
//func work_2() {
//	data := <-piping
//}
//
//func work_3() {
//	data := <-piping
//}
//
//func work_4() {
//	data := <-piping
//}
