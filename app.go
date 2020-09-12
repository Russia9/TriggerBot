package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"triggerbot/lib"
)

var logger = logrus.New()

func main() {
	// Logger init
	logger.Out = os.Stdout
	logger.Info("Initializing ChatWars Broker")

	// Change logger log level
	switch os.Getenv("CWBR_LOGLEVEL") {
	case "TRACE":
		logger.SetLevel(logrus.TraceLevel)
		break
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
		break
	case "INFO":
		logger.SetLevel(logrus.InfoLevel)
		break
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
		break
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
		break
	case "FATAL":
		logger.SetLevel(logrus.FatalLevel)
		break
	case "PANIC":
		logger.SetLevel(logrus.PanicLevel)
		break
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// Sentry init
	logger.Debug("Initializing Sentry")
	SentryDSN := lib.GetEnv("CWBR_SENTRY_DSN", "")
	SentryEnvironment := lib.GetEnv("CWBR_ENVIRONMENT", "production")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         SentryDSN,
		Environment: SentryEnvironment,
	})
	if err != nil {
		logger.Warn("Sentry init error: ", err.Error())
	}
	defer sentry.Flush(2 * time.Second)

	// TODO: Bot init
}
