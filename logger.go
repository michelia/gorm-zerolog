package zapgorm

import (
	"go.uber.org/zap"
)

// FromZap create logger object for *gorm.DB from *zap.Logger
func FromZap(zap *zap.Logger) *Logger {
	return &Logger{
		zap: zap,
	}
}

// Logger is an alternative implementation of *gorm.Logger
type Logger struct {
	zap *zap.Logger
}

// Print passes arguments to Println
func (l *Logger) Print(values ...interface{}) {
	l.Println(values)
}

// Println format & print log
func (l *Logger) Println(values []interface{}) {
	l.zap.Info("gorm", createLog(values).toZapFields()...)
}