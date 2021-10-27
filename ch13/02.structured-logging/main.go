package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var encoderCfg = zapcore.EncoderConfig{
	MessageKey: "msg",
	NameKey:    "name",

	LevelKey:    "level",
	EncodeLevel: zapcore.LowercaseLevelEncoder,

	CallerKey:    "caller",
	EncodeCaller: zapcore.ShortCallerEncoder,

	// TimeKey:    "time",
	// EncodeTime: zapcore.ISO8601TimeEncoder,
}

func main() {
	fmt.Print("\n== JSON Structured Logging Example ==\n")
	jsonStructuredLoggingExample()

	fmt.Print("\n== Console Structured Logging Example ==\n")
	consoleStructuredLoggingExample()

	fmt.Print("\n== Multiple Outputs Structured Logging Example ==\n")
	multipleStructuredLoggingOutputsExample()
}

func jsonStructuredLoggingExample() {
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		),
		zap.AddCaller(),
		zap.Fields(
			zap.String("version", runtime.Version()),
		),
	)
	defer func() { _ = zl.Sync() }()

	example := zl.Named("example")
	example.Debug("test debug message")
	example.Info("test info message")
}

func consoleStructuredLoggingExample() {
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zapcore.InfoLevel,
		),
	)
	defer func() { _ = zl.Sync() }()

	console := zl.Named("[console]")
	console.Info("this is logged by the logger")
	console.Debug("this is below the logger's threshold and won't log")
	console.Error("this is also logged by the logger")
}

func multipleStructuredLoggingOutputsExample() {
	logFile := new(bytes.Buffer)
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(zapcore.AddSync(logFile)),
			zapcore.InfoLevel,
		),
	)
	defer func() { _ = zl.Sync() }()

	zl.Debug("this is below the logger's threshold and won't log")
	zl.Error("this is logged by the logger")

	zl = zl.WithOptions(
		zap.WrapCore(
			func(c zapcore.Core) zapcore.Core {
				ucEncoderCfg := encoderCfg
				ucEncoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
				return zapcore.NewTee(
					c,
					zapcore.NewCore(
						zapcore.NewConsoleEncoder(ucEncoderCfg),
						zapcore.Lock(os.Stdout),
						zapcore.DebugLevel,
					),
				)
			},
		),
	)

	fmt.Println("standard output:")
	zl.Debug("this is only logged as console encoding")
	zl.Info("this is logged as console encoding and JSON")

	fmt.Print("\nlog file contents:\n", logFile.String())
}
