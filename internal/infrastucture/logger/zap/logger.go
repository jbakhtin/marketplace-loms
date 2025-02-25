package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ToDo: рефакторинг
// ToDo: вынести объект Writer в отдельные реализации, что бы можно было подключать разные хранилища логов
// ToDo: разобраться почему некорректно работает форматирования вывода
// ToDo: разобраться почему файл обновляется только после остановки приложения

const (
	DevelopmentEnvironment = "development"
	ProductionEnvironment  = "production"
)

type Config interface {
	GetAppEnv() string
	GetLoggerFileDirectory() string
	GetLoggerFileMaxSize() int
	GetLoggerFileMaxBackups() int
	GetLoggerFileMaxAge() int
	GetLoggerFileCompress() bool
}

type Logger struct {
	zap.Logger
}

func NewLogger(cfg Config) (Logger, error) {
	return Logger{
		Logger: *zap.New(zapcore.NewTee()),
	}, nil
}

func (l Logger) Debug(msg string, fields ...any) {
	l.Logger.Debug(msg, zap.Any("args", fields))
}

func (l Logger) Info(msg string, fields ...any) {
	l.Logger.Info(msg, zap.Any("args", fields))
}

func (l Logger) Warn(msg string, fields ...any) {
	l.Logger.Warn(msg, zap.Any("args", fields))
}

func (l Logger) Error(msg string, fields ...any) {
	l.Logger.Error(msg, zap.Any("args", fields))
}

func (l Logger) Fatal(msg string, fields ...any) {
	l.Logger.Fatal(msg, zap.Any("args", fields))
}
