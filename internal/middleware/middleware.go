package middleware

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func ApiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"body":   string(bodyBytes),
		}).Info("incoming request")

		next.ServeHTTP(w, r)
	})
}
