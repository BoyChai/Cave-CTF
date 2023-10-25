package log

import (
	"Cave-CTF/config"
	"fmt"
	"os"
	"time"
)

func NewLog(level Level, code Code, err error, msg string) {
	data := info{
		Is:    true,
		Level: level,
		Code:  code,
		Err:   err,
		Msg:   msg,
		Date:  time.Now(),
	}

	piping <- data
}

// 输出至前端
func outDisplay(str string) {
	fmt.Printf(str)
}

// 输出至日志文件
func outFile(str string) {
	checkChangeLogfile()
	_, err := logFile.WriteString(str)
	if err != nil {
		fmt.Printf("Log:追加日志内容出现错误:%s", err.Error())
		return
	}
}

func openLogFile() {
	var err error
	logFile, err = os.OpenFile(config.Config.Log.Path+"/"+logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Printf("Log:打开日志文件出现错误:%s", err.Error())
		return
	}

}

func checkChangeLogfile() {
	now := time.Now()
	if now.YearDay() == day {
		return
	}
	logFile.Close()
	postFix := now.Add(-24 * time.Hour).Format("20060102") //昨天的日期
	if err := os.Rename(logFileName, logFileName+"."+postFix); err != nil {
		fmt.Printf("Log:将日期后缀 %s 追加到日志文件 %s 失败: %v\n", postFix, logFileName, err)
		return
	}
	var err error
	logFile, err = os.OpenFile(config.Config.Log.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Printf("Log:打开日志%s错误:%s", logFileName, err.Error())
	}
	day = now.YearDay()
}
