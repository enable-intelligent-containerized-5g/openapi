package models

type LossConnectivityCfg struct {
	// indicating a time in seconds.
	MaxDetectionTime int32 `json:"maxDetectionTime,omitempty" yaml:"maxDetectionTime" bson:"maxDetectionTime,omitempty"`
}
