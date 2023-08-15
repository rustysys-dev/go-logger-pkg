package main

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rustysys-dev/go-logger-pkg/logger"
	"github.com/rustysys-dev/go-logger-pkg/logger/rsdzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger(f io.Writer) zerolog.Logger {
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

	lgr := zerolog.New(f).With().Timestamp().Caller().Logger()
	return lgr
}

type User struct {
	Name string
	Age  int
}

// this is a comment on the main function
func main() {
	zapconfig := rsdzap.NewConfig()
	zapconfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	rsdzap.InitLogger(zapconfig)
	rsdzap.Raw().With("Name", "Tom").Debug("hello")

	lgr := newLogger(os.Stdout)
	lgr.Debug().Str("Name", "Tom").Msg("hello")

	nlgr := logger.NewLogger(os.Stdout)
	nlgr = nlgr.AddItem("MyKey", 123)
	nlgr = nlgr.AddItem("MyKey2", "world")
	nlgr = nlgr.AddItem("userObj", User{Name: "rusty", Age: 99})
	nlgr.Debug("hello")
}
