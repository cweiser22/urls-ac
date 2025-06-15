package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var CacheRequests = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "cache_requests_total",
		Help: "Total number of cache requests labeled by result",
	},
	[]string{"result"}, // result = "hit" or "miss"
)
