package main

import (
	"net/http"
	"net/http/httptest"
	"os"

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
}

func main() {
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		),
	)
	defer func() { _ = zl.Sync() }()

	ts := httptest.NewServer(
		WideEventLog(zl, http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello!"))
			},
		)),
	)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/test")
	if err != nil {
		zl.Fatal(err.Error())
	}
	defer resp.Body.Close()
}
