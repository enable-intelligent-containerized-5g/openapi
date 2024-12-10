package pcm_models

import "time"

type NwdafPacketCaptureMetric struct {
	CpuUsage  float64    `json:"cpu_usage" yaml:"cpu_usage" bson:"cpu_usage"`
	MemUsage  float64    `json:"mem_usage" yaml:"mem_usage" bson:"mem_usage"`
	CpuLimit  float64    `json:"cpu_limit" yaml:"cpu_limit" bson:"cpu_limit"`
	MemLimit  float64    `json:"mem_limit" yaml:"mem_limit" bson:"mem_limit"`
	Timestamp *time.Time `json:"timestamp" yaml:"timestamp" bson:"timestamp"`
}
