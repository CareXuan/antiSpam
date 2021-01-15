package base

import (
	"log"
	"os"
)

var Xlog = XLog{}

type XLog struct {
	InfoInit    *log.Logger
	WarningInit *log.Logger
	ErrorInit   *log.Logger
}

func initXLog(access string, error string) {
	accessFile, err := os.OpenFile(access,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	errorFile, err := os.OpenFile(error,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Xlog.InfoInit = log.New(accessFile,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Xlog.WarningInit = log.New(accessFile,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Xlog.ErrorInit = log.New(errorFile,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(content interface{}) {
	Xlog.InfoInit.Print(content)
}

func Warning(content interface{}) {
	Xlog.WarningInit.Print(content)
}

func Error(content interface{}) {
	Xlog.ErrorInit.Print(content)
}
