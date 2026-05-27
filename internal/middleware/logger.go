package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Method: ", r.Method, "URL: ", r.URL.Path, "Time: ", time.Since(start))
		next.ServeHTTP(w, r)
	})
}
