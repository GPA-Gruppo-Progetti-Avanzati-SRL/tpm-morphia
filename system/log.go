package system

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

const DefaultLogSourceField = "source"
const DefaultLogMessageField = "message"
const DefaultLoggerLevel = "Debug"

type LoggerConfig struct {
	LogLevel string
}

var logger log.Logger = InitializeLogger(LoggerConfig{
	LogLevel: DefaultLoggerLevel,
})

func GetLogger() log.Logger {
	return logger
}

func SetLogger(aLogger log.Logger) {
	logger = aLogger
}

func Fatal(err error) {
	panic(err)
}

/*
 * L'ordine di 'creazione' è importante. Il DefaultCaller va messo dopo la valorizzazione del Level.
 * Se si cambia configurazione conviene crearne uno nuovo oppure mantenere un sottostante nel caso di file aperti o cose del genere.
 */
func InitializeLogger(loggerConfig LoggerConfig) log.Logger {
	aLogger := log.NewLogfmtLogger(os.Stderr)
	aLogger = log.With(aLogger, "ts", log.DefaultTimestampUTC)

	switch loggerConfig.LogLevel {
	case "Debug":
		aLogger = level.NewFilter(aLogger, level.AllowDebug())
	case "Info":
		aLogger = level.NewFilter(aLogger, level.AllowInfo())
	case "Warn":
		aLogger = level.NewFilter(aLogger, level.AllowWarn())
	case "Error":
		aLogger = level.NewFilter(aLogger, level.AllowError())
	}

	aLogger = log.With(aLogger, "caller", log.DefaultCaller)
	_ = level.Info(aLogger).Log(DefaultLogSourceField, "system", DefaultLogMessageField, "Logger Initialized", "LogLevel", loggerConfig.LogLevel)

	return aLogger
}
