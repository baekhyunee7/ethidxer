package ethidxer

import (
	"fmt"
)

type Logger interface {
	Infof(msg string, args ...any)
	Errorf(msg string, args ...any)
	Warnf(msg string, args ...any)
}

type defaultLogger struct{}

func NewDefaultLogger() *defaultLogger {
	return &defaultLogger{}
}

func (a *defaultLogger) Infof(msg string, args ...any) {
	fmt.Printf(msg+"\n", args...)
}

func (a *defaultLogger) Errorf(msg string, args ...any) {
	fmt.Printf(msg+"\n", args...)
}

func (a *defaultLogger) Warnf(msg string, args ...any) {
	fmt.Printf(msg+"\n", args...)
}
