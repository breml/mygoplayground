package main

import (
	"flag"
	"io"
	"log"
	"os"
)

var (
	logLevel  = flag.Int("loglevel", 1, "0 = no output, 1 = standard log, 2 = verbose log")
	logTarget = flag.String("logtarget", "/dev/stdout", "log target")

	logTargetWriter io.Writer

	logger      *log.Logger
	debugLogger *log.Logger

	emptyLogf = func(fmt string, args ...interface{}) {}
	stdLogf   = func(fmt string, args ...interface{}) { logger.Printf(fmt, args...) }
	debugLogf = func(fmt string, args ...interface{}) { debugLogger.Printf(fmt, args...) }

	logf   = stdLogf
	debugf = emptyLogf
)

func main() {
	var err error

	flag.Parse()

	_, err = os.Stat(*logTarget)
	if os.IsExist(err) {
		logTargetWriter, err = os.OpenFile(*logTarget, os.O_WRONLY, 0x600)
		if err != nil {
			log.Fatalln("fatal: unable to open log target", err)
		}
	} else {
		logTargetWriter, err = os.Create(*logTarget)
		if err != nil {
			log.Fatalln("fatal: unable to create log target", err)
		}
	}

	logger = log.New(logTargetWriter, "LOG   ", log.LstdFlags)
	debugLogger = log.New(logTargetWriter, "DEBUG ", log.LstdFlags)

	switch *logLevel {
	case 0:
		logf = emptyLogf
	case 2:
		debugf = debugLogf
	default:
	}

	logf("log output %s, %d", "string1", 10)
	debugf("debug output %s, %d", "string1", 10)
}
