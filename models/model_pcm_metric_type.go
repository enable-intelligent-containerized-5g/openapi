package models

type MetricType string

const (
	MetricType_CPU_USAGE            MetricType = "cpu-usage"
	MetricType_MEMORY_USAGE         MetricType = "mem-usage"
	MetricType_CPU_USAGE_AVERAGE    MetricType = "cpu-average"
	MetricType_MEMORY_USAGE_AVERAGE MetricType = "mem-average"
	MetricType_CPU_LIMIT            MetricType = "cpu-limit"
	MetricType_MEMORY_LIMIT         MetricType = "men-limit"
	MetricType_CPU_REQUEST          MetricType = "cpu-request"
	MetricType_MEMORY_REQUEST       MetricType = "men-request"
	MetricType_POD_STATUS           MetricType = "pod-status"
	MetricType_RUNNING_POD          MetricType = "running-pod"
)
