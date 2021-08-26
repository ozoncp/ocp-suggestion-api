package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter prometheus.Counter
	updateCounter prometheus.Counter
	removeCounter prometheus.Counter
)

func RegisterMetrics() {
	createCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_suggestion_api_create_count_total",
		Help: "The total number of created suggestions",
	})

	updateCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_suggestion_api_update_count_total",
		Help: "The total number of updated suggestions",
	})

	removeCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_suggestion_api_remove_count_total",
		Help: "The total number of removed suggestions",
	})
}

func CreateCounterInc() {
	createCounter.Inc()
}

func UpdateCounterInc() {
	updateCounter.Inc()
}

func RemoveCounterInc() {
	removeCounter.Inc()
}
