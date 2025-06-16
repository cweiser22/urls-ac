package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var CacheRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "redirect_cache_hits_total",
		Help: "Number of cache hits and misses for URL redirection",
	},
	[]string{"result"}, // result = "hit" or "miss"
)
