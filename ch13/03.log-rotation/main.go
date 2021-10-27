package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var encoderCfg = zapcore.EncoderConfig{
	MessageKey: "msg",
	NameKey:    "name",

	LevelKey:    "level",
	EncodeLevel: zapcore.LowercaseLevelEncoder,

	CallerKey:    "caller",
	EncodeCaller: zapcore.ShortCallerEncoder,
}

func main() {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.AddSync(
				&lumberjack.Logger{
					Filename:   filepath.Join(tempDir, "debug.log"),
					Compress:   true,
					LocalTime:  true,
					MaxAge:     7,
					MaxBackups: 5,
					MaxSize:    100,
				},
			),
			zapcore.DebugLevel,
		),
	)
	defer func() { _ = zl.Sync() }()

	zl.Debug("debug message written to the log file")
}
