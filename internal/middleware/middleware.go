package middleware

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"user-service/pkg"
)

func ApiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)

			if err != nil {
				logrus.WithError(err).Error("Failed to read request body")
				pkg.BadRequest(err, w)
			}

			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
			logrus.WithFields(logrus.Fields{
				"method":     r.Method,
				"path":       r.URL.Path,
				"user_agent": r.UserAgent(),
				"ip":         r.RemoteAddr,
				"body":       string(bodyBytes),
			}).Info("incoming request")
		} else {
			logrus.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("incoming request")
		}

		next.ServeHTTP(w, r)
	})
}
