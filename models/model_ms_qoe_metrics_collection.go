package models

// Contains the Media Streaming QoE metrics information collected for an UE Application via AF.
type MsQoeMetricsCollection struct {
	MsQoeMetrics []string `json:"msQoeMetrics" yaml:"msQoeMetrics" bson:"msQoeMetrics,omitempty"`
}
