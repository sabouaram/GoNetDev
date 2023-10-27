package Utils

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func init() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(": 8080", nil))
	}()
}
func NewPromCounter(metric string, help_msg string, labels []string) *prometheus.CounterVec {

	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metric,
			Help: help_msg,
		},
		labels,
	)

}

func NewLabel(lblkey string, lblvalue string) prometheus.Labels {
	return prometheus.Labels{lblkey: lblvalue}
}

func Register(c prometheus.Collector) {
	prometheus.MustRegister(c)
}
