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

func withLevel(event *zerolog.Event, msg ...interface{}) {
	event.Msg(fmt.Sprint(msg...))
}

func withLevelf(event *zerolog.Event, format string, v ...interface{}) {
	event.Msgf(format, v...)
}

func withLevelWhenError(event *zerolog.Event, err error, msg ...interface{}) {
	if err != nil {
		event.Err(err).Msg(fmt.Sprint(msg...))
	}
}

func withLevelWhenErrorf(event *zerolog.Event, err error, format string, v ...interface{}) {
	if err != nil {
		event.Err(err).Msgf(format, v)
	}
}

func Debug(msg ...interface{}) {
	withLevel(log.Debug(), msg...)
}

func Info(msg ...interface{}) {
	withLevel(log.Info(), msg...)
}

func Warn(err error, msg ...interface{}) {
	withLevelWhenError(log.Warn(), err, msg...)
}

func Fatal(err error, msg ...interface{}) {
	withLevelWhenError(log.Fatal(), err, msg...)
}

func Debugf(format string, v ...interface{}) {
	withLevelf(log.Debug(), format, v...)
}

func Infof(format string, v ...interface{}) {
	withLevelf(log.Info(), format, v...)
}

func Fatalf(err error, format string, v ...interface{}) {
	withLevelWhenErrorf(log.Fatal(), err, format, v...)
}

func Warnf(err error, format string, v ...interface{}) {
	withLevelWhenErrorf(log.Warn(), err, format, v...)
}
