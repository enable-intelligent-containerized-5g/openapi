package models

type UdmEeReachabilityForDataConfiguration struct {
	ReportCfg ReachabilityForDataReportConfig `json:"reportCfg" yaml:"reportCfg" bson:"reportCfg,omitempty"`
	// indicating a time in seconds.
	MinInterval int32 `json:"minInterval,omitempty" yaml:"minInterval" bson:"minInterval,omitempty"`
}
