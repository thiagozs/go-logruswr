package logruswr

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogWrapper struct {
	log      *logrus.Logger
	exitFunc func(int)
}

func New(opts ...Options) (*LogWrapper, error) {
	params, err := newLogWrapperParams(opts...)
	if err != nil {
		return nil, err
	}

	log := logrus.New()

	// Set formatter
	switch params.GetFormatter() {
	case FormatterText:
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			ForceColors:     true,
		})
	case FormatterJSON:
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}

	// Set log output
	switch params.output {
	case Stdout:
		log.SetOutput(os.Stdout)
	case Stderr:
		log.SetOutput(os.Stderr)
	case File:
		if params.GetLogFilePath() != "" {
			log.SetOutput(&lumberjack.Logger{
				Filename:   params.GetLogFilePath(),
				MaxSize:    params.GetMaxLogSize(), // megabytes
				MaxBackups: params.GetMaxBackups(),
				MaxAge:     params.GetMaxAge(), // days
				Compress:   params.GetCompressLogs(),
			})
		} else {
			log.SetOutput(os.Stdout)
		}
	}

	// Set log level
	log.SetLevel(logrus.Level(params.GetLevel()))

	// Add hooks
	for _, hook := range params.hooks {
		log.AddHook(hook)
	}

	return &LogWrapper{
		log:      log,
		exitFunc: os.Exit,
	}, nil
}

func (l *LogWrapper) WithField(key string, value interface{}) *Entry {
	return l.log.WithField(key, value)
}

func (l *LogWrapper) WithFields(fields Fields) *Entry {
	return l.log.WithFields(logrus.Fields(fields))
}

func (l *LogWrapper) WithError(err error) *Entry {
	return l.log.WithError(err)
}

func (l *LogWrapper) WithContext(ctx context.Context) *Entry {
	return l.log.WithContext(ctx)
}

func (l *LogWrapper) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *LogWrapper) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *LogWrapper) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *LogWrapper) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *LogWrapper) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *LogWrapper) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *LogWrapper) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *LogWrapper) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
	l.exitFunc(1) // exit

}

func (l *LogWrapper) Panic(args ...interface{}) {
	l.log.Panic(args...)
}

func (l *LogWrapper) Trace(args ...interface{}) {
	l.log.Trace(args...)
}

func (l *LogWrapper) AddHook(hook Hook) {
	l.log.AddHook(hook)
}

func (l *LogWrapper) SetLevel(level Level) {
	l.log.SetLevel(logrus.Level(level))
}

func (l *LogWrapper) SetReportCaller(b bool) {
	l.log.SetReportCaller(b)
}

func (l *LogWrapper) SetExitFunc(exitFunc func(int)) {
	l.exitFunc = exitFunc
}
