package log

import (
	"go.uber.org/zap"

	"github.com/masibw/go-docker-template/config"
)

// logger is the interface of the logger.
type logger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

// Logger is the reality called by the program
var Logger logger

// Structure for zap
type zapLogger struct {
	logger *zap.Logger
}

func init() {
	var zapLog *zap.Logger
	if config.IsProd() {
		zapLog, _ = zap.NewProduction(zap.AddCallerSkip(1))
	} else {
		zapLog, _ = zap.NewDevelopment(zap.AddCaller(), zap.AddCallerSkip(1))
	}
	Logger = &zapLogger{
		logger: zapLog,
	}
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	defer func() {
		_ = l.logger.Sync()
	}()
	l.logger.Sugar().Debugf(template, args...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	defer func() {
		_ = l.logger.Sync()
	}()
	l.logger.Sugar().Infof(template, args...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	defer func() {
		_ = l.logger.Sync()
	}()
	l.logger.Sugar().Warnf(template, args...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	defer func() {
		_ = l.logger.Sync()
	}()
	l.logger.Sugar().Errorf(template, args...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	defer func() {
		_ = l.logger.Sync()
	}()
	l.logger.Sugar().Fatalf(template, args...)
}
