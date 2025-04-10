package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Технические
	RequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Количество HTTP-запросов",
		},
		[]string{"method", "path"},
	)

	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Время обработки запросов",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	// Бизнесовые
	PVZCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "pvz_created_total",
		Help: "Количество созданных ПВЗ",
	})

	ReceptionsCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "receptions_created_total",
		Help: "Количество созданных приёмок",
	})

	ProductsCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_created_total",
		Help: "Количество добавленных товаров",
	})
)
