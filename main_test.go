package main

import (
	"os"
	"testing"

	"github.com/rustysys-dev/go-logger-pkg/logger"
)

var devnull, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0)

func BenchmarkNewLoggerObject(b *testing.B) {
	runloop(10000, func() {
		// nlgr := logger.NewLogger(&bytes.Buffer{})
		nlgr := logger.NewLogger(devnull)
		nlgr = nlgr.AddItem("userObj", User{Name: "rusty", Age: 99})
		nlgr.Debug("hello")
	})
}

func BenchmarkNewLoggerString(b *testing.B) {
	runloop(10000, func() {
		// nlgr := logger.NewLogger(&bytes.Buffer{})
		nlgr := logger.NewLogger(devnull)
		nlgr = nlgr.AddItem("MyKey2", "world")
		nlgr.Debug("hello")
	})
}

func BenchmarkNewLoggerInt(b *testing.B) {
	runloop(10000, func() {
		// nlgr := logger.NewLogger(&bytes.Buffer{})
		nlgr := logger.NewLogger(devnull)
		nlgr = nlgr.AddItem("MyKey", 123)
		nlgr.Debug("hello")
	})
}

func BenchmarkZerologObject(b *testing.B) {
	runloop(10000, func() {
		// lgr := newLogger(&bytes.Buffer{})
		lgr := newLogger(devnull)
		lgr.Debug().Interface("userObj", User{Name: "rusty", Age: 99}).Msg("hello")
	})
}

func BenchmarkZerologString(b *testing.B) {
	runloop(10000, func() {
		// lgr := newLogger(&bytes.Buffer{})
		lgr := newLogger(devnull)
		lgr.Debug().Str("MyKey2", "world").Msg("hello")
	})
}

func BenchmarkZerologInt(b *testing.B) {
	runloop(10000, func() {
		// lgr := newLogger(&bytes.Buffer{})
		lgr := newLogger(devnull)
		lgr.Debug().Int("MyKey", 123).Msg("hello")
	})
}

func runloop(times int, fn func()) {
	for i := 0; i < times; i++ {
		fn()
	}
}
