package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Labels = []string{"amount", "x_trace_id", "event_type"}

func NewLabels(amount, traceID, eventType string) prometheus.Labels {
	return prometheus.Labels{
		"amount":     amount,
		"x_trace_id": traceID,
		"event_type": eventType,
	}
}

func Record(metric *prometheus.SummaryVec, labels prometheus.Labels) {
	timer := prometheus.NewTimer(metric.With(labels))
	timer.ObserveDuration()
}
