package xlog

import (
	"fmt"
	"os"
	"time"

	isatty "github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var DefaultTimeFieldFormat = time.RFC3339Nano
var SimpleTimeFieldFormat = "2006-01-02 15:04:05.000"

func init() {
	if isatty.IsTerminal(os.Stderr.Fd()) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	switch os.Getenv("XLOG_LEVEL") {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "", "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "NOLEVEL":
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	case "DISABLED":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
	zerolog.TimeFieldFormat = DefaultTimeFieldFormat
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

func WarnMsg(msg ...interface{}) {
	withLevel(log.Warn(), msg...)
}

func Error(err error, msg ...interface{}) {
	withLevelWhenError(log.Error(), err, msg...)
}

func ErrorMsg(msg ...interface{}) {
	withLevel(log.Error(), msg...)
}

func Fatal(err error, msg ...interface{}) {
	withLevelWhenError(log.Fatal(), err, msg...)
}

func FatalMsg(msg ...interface{}) {
	withLevel(log.Fatal(), msg...)
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

func Errorf(err error, format string, v ...interface{}) {
	withLevelWhenErrorf(log.Error(), err, format, v...)
}

func WarnMsgf(format string, msg ...interface{}) {
	withLevelf(log.Warn(), format, msg...)
}

func ErrorMsgf(format string, msg ...interface{}) {
	withLevelf(log.Error(), format, msg...)
}

func FatalMsgf(format string, msg ...interface{}) {
	withLevelf(log.Fatal(), format, msg...)
}
