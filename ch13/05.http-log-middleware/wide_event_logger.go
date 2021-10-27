package main

import (
	"net"
	"net/http"

	"go.uber.org/zap"
)

type wideResponseWriter struct {
	http.ResponseWriter
	length, status int
}

func (w *wideResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.status = status
}

func (w *wideResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.length += n

	if w.status == 0 {
		// 200 OK inferred on first Write if status is not yet set
		w.status = http.StatusOK
	}

	return n, err
}

func WideEventLog(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wideWriter := &wideResponseWriter{ResponseWriter: w}

			next.ServeHTTP(wideWriter, r)

			addr, _, _ := net.SplitHostPort(r.RemoteAddr)
			logger.Info("example wide event",
				zap.Int("status_code", wideWriter.status),
				zap.Int("response_length", wideWriter.length),
				zap.Int64("content_length", r.ContentLength),
				zap.String("method", r.Method),
				zap.String("proto", r.Proto),
				zap.String("remote_addr", addr),
				zap.String("uri", r.RequestURI),
				zap.String("user_agent", r.UserAgent()),
			)
		},
	)
}
