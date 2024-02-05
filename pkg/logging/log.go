package logging

import (
	"fmt"
	"github.com/marsli9945/go-pkg/pkg/file"
	"github.com/marsli9945/go-pkg/pkg/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

const (
	green  = "\033[97;42m"
	white  = "\033[90;47m"
	yellow = "\033[90;43m"
	red    = "\033[97;41m"
	//blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

var (
	DefaultCallerDepth = 2

	logger            *log.Logger
	errLogger         *log.Logger
	schedulerLogger   *log.Logger
	logPrefix         = ""
	levelFlags        = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	levelColors       = []string{white, green, yellow, red, magenta}
	filePath          = "runtime/logs/"
	fileName          = "app.log"
	errFileName       = "error.log"
	schedulerFileName = "scheduler.log"
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// setup initialize the log instance
func setup() {
	var err error
	var F *os.File
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
	logger = log.New(io.MultiWriter(F, os.Stdout), "", 0)

	F, err = file.MustOpen(errFileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
	errLogger = log.New(io.MultiWriter(F, os.Stdout), "", 0)

	F, err = file.MustOpen(schedulerFileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
	schedulerLogger = log.New(F, "", 0)
}

// getLogger init and get logger
func getLogger() *log.Logger {
	if logger == nil {
		setup()
	}
	return logger
}

// getErrorLogger init and get errLogger
func getErrorLogger() *log.Logger {
	if errLogger == nil {
		setup()
	}
	return errLogger
}

// getSchedulerLogger init and get schedulerLogger
func getSchedulerLogger() *log.Logger {
	if schedulerLogger == nil {
		setup()
	}
	return schedulerLogger
}

// Info output logs at info level
func Info(v ...any) {
	setPrefix(INFO)
	getLogger().Println(v...)
}

// InfoF output logs at infof level
func InfoF(format string, v ...any) {
	setPrefix(INFO)
	getLogger().Printf(format, v...)
}

func InfoJson(v any) {
	setPrefix(INFO)
	getLogger().Println(util.JsonMarshal(v))
}

func InfoJsonFormat(v any) {
	setPrefix(INFO)
	getLogger().Println("\r\n" + util.PrettyString(util.JsonMarshal(v)))
}

// Warn output logs at warn level
func Warn(v ...any) {
	setPrefix(WARNING)
	getLogger().Println(v...)
}

// Error output logs at error level
func Error(v ...any) {
	setPrefix(ERROR)
	getErrorLogger().Println(v...)
}

// ErrorF output logs at error level
func ErrorF(format string, v ...any) {
	setPrefix(ERROR)
	getErrorLogger().Printf(format, v...)
}

// Fatal output logs at fatal level
func Fatal(v ...any) {
	setPrefix(FATAL)
	getErrorLogger().Fatalln(v...)
}

func SchedulerInfoF(format string, v ...any) {
	setPrefix(INFO)
	getSchedulerLogger().Printf(format, v...)
}

func SchedulerErrorF(format string, v ...any) {
	setPrefix(ERROR)
	getSchedulerLogger().Printf(format, v...)
}

func getTime() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, runtimeFile, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%sMARS%s] %s |%s %-5s %s| %s:%-3d ", cyan, reset, getTime(), levelColors[level], levelFlags[level], reset, filepath.Base(runtimeFile), line)
	} else {
		logPrefix = fmt.Sprintf("[%sMARS%s] %s |%s %-5s %s| ", cyan, reset, getTime(), levelColors[level], levelFlags[level], reset)
	}
	getLogger().SetPrefix(logPrefix)
	getErrorLogger().SetPrefix(logPrefix)
	getSchedulerLogger().SetPrefix(logPrefix)
}
