package prometheus

import (
	"sort"
	"time"

	prommodel "github.com/prometheus/common/model"
)

// TimeDurationToPromStr converts from std duration to prom string duration.
func TimeDurationToPromStr(t time.Duration) string {
	return prommodel.Duration(t).String()
}

// LabelsToPromFilter converts a map of labels to a Prometheus filter string.
func LabelsToPromFilter(labels map[string]string) string {
	metricFilters := prommodel.LabelSet{}
	for k, v := range labels {
		metricFilters[prommodel.LabelName(k)] = prommodel.LabelValue(v)
	}

	return metricFilters.String()
}

// LabelsToPromMatch converts a map of labels to a Prometheus filter string.
func LabelsToPromGroup(labels map[string]string) string {
	metricGroup := prommodel.LabelNames{}
	for k, _ := range labels {
		metricGroup = append(metricGroup, prommodel.LabelName(k))
	}

	sort.Sort(metricGroup)
	return metricGroup.String()
}
