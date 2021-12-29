package loger

import (
	"Learn-CasaOS/pkg/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	file2 "Learn-CasaOS/pkg/utils/file"
)

// 定义一个int的别名
type Level int

type Olog interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
}

type oLog struct {
}

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

// iota在const关键字出现时将被重置为0(const内部的第一行之前)，
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

// 日志初始化
func LogSetup() {
	var err error
	// filePath := fmt.Sprintf("%s", config.AppInfo.LogSavePath)
	filePath := config.AppInfo.LogSavePath
	fileName := fmt.Sprintf("%s%s.%s", config.AppInfo.LogSaveName, time.Now().Format(config.AppInfo.DateStrFormat), config.AppInfo.LogFileExt)
	F, err = file2.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func (o *oLog) Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func (o *oLog) Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func (o *oLog) Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v)
}

func (o *oLog) Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func (o *oLog) Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

func setPrefix(lever Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[lever], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[lever])
	}

	logger.SetPrefix(logPrefix)
}

func NewOLoger() Olog {
	return &oLog{}
}
