package log

import (
	"os"
	"time"
)

// 时间
var day int

// 日志文件
var logFile *os.File

// 日志名称
var logFileName string

// info 错误信息
type info struct {
	Is    bool
	Level Level
	Code  Code
	Err   error
	Msg   string
	Date  time.Time
}

// Level 错误级别
type Level int

const (
	DEBUG    Level = 0
	INFO     Level = 1
	WARNING  Level = 2
	ERROR    Level = 3
	CRITICAL Level = 4
)

// Code 错误码
type Code int
