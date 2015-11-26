package main

import (
	"flag"
	"log"
	"os"
)

var (
	logLevel = flag.Int("loglevel", 1, "0 = no output, 1 = standard log, 2 = verbose log")

	logger      = log.New(os.Stdout, "LOG   ", log.LstdFlags)
	debugLogger = log.New(os.Stdout, "DEBUG ", log.LstdFlags)

	emptyLogf = func(fmt string, args ...interface{}) {}
	stdLogf   = func(fmt string, args ...interface{}) { logger.Printf(fmt, args...) }
	debugLogf = func(fmt string, args ...interface{}) { debugLogger.Printf(fmt, args...) }

	logf   = stdLogf
	debugf = emptyLogf
)

func main() {
	flag.Parse()

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
