package rsdzap

import (
	"bytes"
	"context"
	"net/url"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	gcpDebugString = "DEBUG"
	gcpInfoString  = "INFO"
	gcpErrorString = "ERROR"
	gcpFatalString = "CRITICAL"
)

// All loggers are spawned from this initial core
var lgr *zap.SugaredLogger

func init() {
	InitLogger(NewConfig())
}

// InitLogger sets the initial core logger
func InitLogger(conf *zap.Config) {
	l, e := conf.Build()
	if e != nil {
		panic(e)
	}

	lgr = l.Sugar()
}

type Logger struct {
	*zap.SugaredLogger
}

// Type sets the log type
func (l *Logger) Type(t string) *Logger {
	return &Logger{l.With("type", t)}
}

// Print implementation for various logger interfaces
func (l *Logger) Print(v ...any) {
	l.Warn(v...)
}

// Trace spawns a logger with the user trace
func Trace(c context.Context) *Logger {
	return &Logger{
		lgr.
			With("logging.googleapis.com/trace", c.Value("trace-id-key")).
			With("uri", c.Value("uri-key")),
	}
}

// System spawns a logger with the system trace
func System(c context.Context) *Logger {
	return &Logger{
		lgr.
			With("logging.googleapis.com/trace", "SYSTEM").
			With("uri", c.Value("uri-key")),
	}
}

// System spawns a logger with the system trace
func Raw() *Logger {
	return &Logger{lgr}
}

// func FrontendProxy(c *fiber.Ctx) error {
// 	var v interface{}
// 	err := json.Unmarshal(c.Body(), &v)
// 	if err != nil {
// 		return err
// 	}

// 	lvl := c.Query("level", "info")
// 	if lvl == "" {
// 		lvl = "info"
// 	}

// 	l := WithTrace(c).With("level", lvl)
// 	switch strings.ToLower(lvl) {
// 	case "panic", "fatal", "error":
// 		l.Error(v)
// 	case "warn", "warning", "info", "debug", "trace":
// 		l.Info(v)
// 	}
// 	return nil
// }

// https://cloud.google.com/logging/docs/structured-logging#special-payload-fields
func NewConfig() *zap.Config {
	return &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "textPayload",
			LevelKey:      "severity",
			TimeKey:       "time",
			NameKey:       "",
			CallerKey:     "logging.googleapis.com/sourceLocation",
			StacktraceKey: "",
			// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity
			EncodeLevel: levelEncoder,
			// https://cloud.google.com/logging/docs/agent/logging/configuration#timestamp-processing
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogEntrySourceLocation
			// FIXME check on this
			EncodeCaller: zapcore.FullCallerEncoder,
			EncodeName:   zapcore.FullNameEncoder,
		},
		InitialFields: map[string]any{},
		OutputPaths:   []string{"stdout"},
	}
}

func levelEncoder(l zapcore.Level, e zapcore.PrimitiveArrayEncoder) {
	e.AppendString(map[zapcore.Level]string{
		zapcore.DebugLevel: gcpDebugString,
		zapcore.InfoLevel:  gcpInfoString,
		zapcore.WarnLevel:  gcpErrorString,
		zapcore.ErrorLevel: gcpErrorString,
		zapcore.PanicLevel: gcpErrorString,
		zapcore.FatalLevel: gcpFatalString,
	}[l])
}

// MemorySink implements zap.Sink by writing all messages to a buffer.
type MemorySink struct {
	Mutex *sync.Mutex
	*bytes.Buffer
}

// Implement Close and Sync as no-ops to satisfy the interface. The Write
// method is provided by the embedded buffer.

func (*MemorySink) Close() error {
	return nil
}

func (*MemorySink) Sync() error {
	return nil
}

func NewTestLogger(s *MemorySink) error {
	if err := zap.RegisterSink("memory", func(*url.URL) (zap.Sink, error) {
		return s, nil
	}); err != nil {
		return err
	}

	conf := NewConfig()
	conf.OutputPaths = append(conf.OutputPaths, "memory://")
	InitLogger(conf)
	return nil
}
