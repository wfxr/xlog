package xlog

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if isatty.IsTerminal(os.Stderr.Fd()) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000"
}

func logWithLevel(event *zerolog.Event, msg ...interface{}) {
	event.Msg(fmt.Sprint(msg...))
}

func logfWithLevel(event *zerolog.Event, format string, v ...interface{}) {
	event.Msgf(format, v...)
}

func logWhenErrWithLevel(event *zerolog.Event, err error, msg ...interface{}) {
	if err != nil {
		event.Err(err).Msg(fmt.Sprint(msg...))
	}
}

func logfWhenErrWithLevel(event *zerolog.Event, err error, format string, v ...interface{}) {
	if err != nil {
		event.Err(err).Msgf(format, v)
	}
}

func LogDebug(msg ...interface{}) {
	logWithLevel(log.Debug(), msg...)
}

func LogInfo(msg ...interface{}) {
	logWithLevel(log.Info(), msg...)
}

func LogWarn(err error, msg ...interface{}) {
	logWhenErrWithLevel(log.Warn(), err, msg...)
}

func LogFatal(err error, msg ...interface{}) {
	logWhenErrWithLevel(log.Fatal(), err, msg...)
}

func LogfDebug(format string, v ...interface{}) {
	logfWithLevel(log.Debug(), format, v...)
}

func LogfInfo(format string, v ...interface{}) {
	logfWithLevel(log.Info(), format, v...)
}

func LogfFatal(err error, format string, v ...interface{}) {
	logfWhenErrWithLevel(log.Fatal(), err, format, v...)
}

func LogfWarn(err error, format string, v ...interface{}) {
	logfWhenErrWithLevel(log.Warn(), err, format, v...)
}
