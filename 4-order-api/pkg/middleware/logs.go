package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)

		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.WithFields(logrus.Fields{
			"status":   wrapper.StatusCode,
			"method":   r.Method,
			"url":      r.URL.Path,
			"duration": time.Since(start),
		}).Info("Request processed")
	})
}
