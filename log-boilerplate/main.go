package main

import (
	"flag"
	"io"
	"log"
	"os"
)

var (
	logLevel      = flag.Int("loglevel", 1, "0 = no output, 1 = standard log, 2 = verbose log")
	logTarget     = flag.String("logtarget", "", "log target, default: stdout")
	logPermission = flag.Int("logperm", 0600, "permission for log file, default in ocatal: 0600")

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

	// Set logTarget
	if *logTarget == "" {
		logger = log.New(os.Stdout, "LOG   ", log.LstdFlags)
		debugLogger = log.New(os.Stdout, "DEBUG ", log.LstdFlags)
	} else {
		var logFile *os.File
		fileMode := os.FileMode(*logPermission)
		_, err = os.Stat(*logTarget)
		if os.IsNotExist(err) {
			logFile, err = os.Create(*logTarget)
			if err != nil {
				log.Fatalln("fatal: unable to create log target", err)
			}

			err = os.Chmod(*logTarget, fileMode)
			if err != nil {
				log.Fatalln("fatal: unable to set file permissions", err)
			}
		} else {
			logFile, err = os.OpenFile(*logTarget, os.O_WRONLY|os.O_APPEND, fileMode)
			if err != nil {
				log.Fatalln("fatal: unable to open log target", err)
			}
		}
		defer logFile.Close()
		logTargetWriter = io.Writer(logFile)

		// Check if logTarget is writable
		_, err = logTargetWriter.Write([]byte(""))
		if err != nil {
			log.Fatalln("fatal: unable to write to log target")
		}

		logger = log.New(logTargetWriter, "LOG   ", log.LstdFlags)
		debugLogger = log.New(logTargetWriter, "DEBUG ", log.LstdFlags)
	}

	// Set log-function based on log level
	switch *logLevel {
	case 0:
		logf = emptyLogf
	case 2:
		debugf = debugLogf
	default:
	}

	// Write example log lines
	logf("log output %s, %d", "string1", 10)
	debugf("debug output %s, %d", "string1", 10)
}
