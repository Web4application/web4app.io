import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

var (
    apiRequests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "api_requests_total",
            Help: "Total number of API requests",
        },
        []string{"endpoint"},
    )
)

func init() {
    prometheus.MustRegister(apiRequests)
}

func IncrementAPIRequest(endpoint string) {
    apiRequests.WithLabelValues(endpoint).Inc()
}

func ExposeMetrics() {
    http.Handle("/metrics", promhttp.Handler())
    go http.ListenAndServe(":5000", 8080)
}
