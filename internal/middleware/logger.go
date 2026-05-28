package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var bodyLog interface{} // ← Changed from string to interface{}
		if r.Body != nil && r.ContentLength > 0 && r.ContentLength < 1024*10 {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				// Try to parse as JSON for nice formatting
				var parsed interface{}
				if json.Unmarshal(bodyBytes, &parsed) == nil {
					bodyLog = parsed // ← Log as object, not string
				} else {
					bodyLog = string(bodyBytes) // Fallback to string
				}
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapper, r)

		slog.Info("HTTP request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", wrapper.statusCode,
			"duration_ms", time.Since(start).Milliseconds(),
			"remote_addr", r.RemoteAddr,
			"body", bodyLog, // ← Now logs as JSON object!
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
