package rsdzero

import (
	"bytes"
	"fmt"
	"io"

	"github.com/rs/zerolog"
	"github.com/rustysys-dev/go-logger-pkg/logger/scheme"
)

type rsdzeroLogger struct {
	l zerolog.Logger
}

func New(f io.Writer) scheme.Logger {
	zerolog.LevelFieldName = "severity"
	zerolog.LevelTraceValue = "DEBUG"
	zerolog.LevelDebugValue = "DEBUG"
	zerolog.LevelInfoValue = "INFO"
	zerolog.LevelWarnValue = "WARNING"
	zerolog.LevelErrorValue = "ERROR"
	zerolog.LevelFatalValue = "CRITICAL"
	zerolog.LevelPanicValue = "CRITICAL"
	zerolog.CallerFieldName = "logging.googleapis.com/sourceLocation"
	zerolog.MessageFieldName = "textPayload"

	return rsdzeroLogger{
		l: zerolog.New(f).With().Timestamp().Caller().Logger(),
	}
}

func (r rsdzeroLogger) AddItem(key string, data any) scheme.Logger {
	return rsdzeroLogger{r.l.Hook(
		zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
			e.Bytes(key, toByteArray[any](data))
		}),
	)}
}

func (r rsdzeroLogger) Debug(msg string, keyval ...any) {
	r.l.Debug().Msg(msg)
}

func toByteArray[T any](data T) []byte {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%+v", data)

	return buffer.Bytes()
}
