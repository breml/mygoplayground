package main

import (
	"flag"
	"io"
	"log"
	"log/syslog"
	"net"
	"net/url"
	"os"
	"strings"
)

var (
	logLevel      = flag.Int("loglevel", 1, "0 = no output, 1 = standard log, 2 = verbose log")
	logTarget     = flag.String("logtarget", "", "log target, if not set it defaults to stdout, accepts URL, e.g. syslog://syslog-host:514, file:///path/to/file.log")
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
	flag.Parse()

	// Set logTarget
	if *logTarget == "" {
		logger = log.New(os.Stdout, "LOG   ", log.LstdFlags)
		debugLogger = log.New(os.Stdout, "DEBUG ", log.LstdFlags)
	} else {
		u, err := url.Parse(*logTarget)
		if err != nil {
			log.Fatalln("fatal: unable to parse logtarget url", err)
		}

		switch {
		case (u.Scheme == "" || u.Scheme == "file") && u.Path != "":

			var logFile *os.File
			fileMode := os.FileMode(*logPermission)
			_, err = os.Stat(u.Path)
			if os.IsNotExist(err) {
				logFile, err = os.Create(u.Path)
				if err != nil {
					log.Fatalln("fatal: unable to create log target", err)
				}

				err = os.Chmod(u.Path, fileMode)
				if err != nil {
					log.Fatalln("fatal: unable to set file permissions", err)
				}
			} else {
				logFile, err = os.OpenFile(u.Path, os.O_WRONLY|os.O_APPEND, fileMode)
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

		case strings.HasPrefix(u.Scheme, "syslog"):
			var syslogLogger *syslog.Writer
			priority := syslog.LOG_NOTICE | syslog.LOG_LOCAL0
			tag := ""

			if u.Host == "" {
				syslogLogger, err = syslog.New(priority, tag)
				if err != nil {
					log.Fatalln("fatal: unable to connect to syslog", *logTarget, "err:", err)
				}
			} else {
				schemeParts := strings.Split(u.Scheme, "+")
				network := "tcp"
				if len(schemeParts) > 1 {
					network = schemeParts[1]
				}

				_, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					u.Host = u.Host + ":514"
				}

				syslogLogger, err = syslog.Dial(network, u.Host, priority, tag)
				if err != nil {
					log.Fatalln("fatal: unable to connect to syslog", *logTarget, "err:", err)
				}
			}
			defer syslogLogger.Close()
			logTargetWriter = io.Writer(syslogLogger)

		default:
			log.Fatalln("fatal: no valid schema:", *logTarget)
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
