package metric

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricRecorder struct {
	timer *prometheus.Timer
}

var Labels = []string{"amount", "x_trace_id", "event_type"}

func NewLabels(amount, traceID, eventType string) prometheus.Labels {
	return prometheus.Labels{
		"amount":     amount,
		"x_trace_id": traceID,
		"event_type": eventType,
	}
}

func NewRecorder() *MetricRecorder {
	return &MetricRecorder{}
}

func (m *MetricRecorder) WithTimer(metric *prometheus.HistogramVec, labels prometheus.Labels) *MetricRecorder {
	m.timer = prometheus.NewTimer(metric.With(labels))

	return m
}

func (m *MetricRecorder) RecordDuration() {
	duration := m.timer.ObserveDuration()
	log.Printf("Send metrics duration=%d", duration)
}
