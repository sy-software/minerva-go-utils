package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogLevel string

const (
	LOG_PANIC LogLevel = "PANIC"
	LOG_FATAL LogLevel = "FATAL"
	LOG_ERROR LogLevel = "ERROR"
	LOG_WARN  LogLevel = "WARN"
	LOG_INFO  LogLevel = "INFO"
	LOG_DEBUG LogLevel = "DEBUG"
	LOG_TRACE LogLevel = "TRACE"
)

// ConfigureLogger prepares the logger at global level
// It can activate logger pretty print (console output)
// for friendly logs during development and set the log level
func ConfigureLogger(level LogLevel, console bool) {
	if console {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	switch level {
	case LOG_PANIC:
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case LOG_FATAL:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case LOG_ERROR:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case LOG_WARN:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case LOG_INFO:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LOG_DEBUG:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case LOG_TRACE:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
