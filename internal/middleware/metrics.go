package middleware

import (
	"net/http"
	"time"

	"github.com/Egorpalan/api-pvz/internal/metrics"
	"github.com/go-chi/chi/v5"
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		routeContext := chi.RouteContext(r.Context())
		path := routeContext.RoutePattern()

		next.ServeHTTP(w, r)

		duration := time.Since(start).Seconds()
		metrics.RequestCount.WithLabelValues(r.Method, path).Inc()
		metrics.RequestDuration.WithLabelValues(r.Method, path).Observe(duration)
	})
}
