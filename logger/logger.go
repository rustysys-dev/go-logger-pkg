package logger

import (
	"io"

	"github.com/rustysys-dev/go-logger-pkg/logger/rsdzero"
	"github.com/rustysys-dev/go-logger-pkg/logger/scheme"
)

func NewLogger(f io.Writer) scheme.Logger {
	return rsdzero.New(f)
}
