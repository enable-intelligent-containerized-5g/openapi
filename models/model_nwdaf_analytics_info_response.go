package models

type NwdaAnalyticsInfoResponse struct {
	EventId      EventId              `json:"eventId" yaml:"eventId" bson:"eventId"`
	Accuracy     NwdafMlModelAccuracy `json:"accuracy" yaml:"accuracy" bson:"accuracy"`
	NfInstanceId string               `json:"nfInstanceId" yaml:"nfInstanceId" bson:"nfInstanceId"`
	NfTypes      NfType               `json:"nfTypes" yaml:"nfTypes" bson:"nfTypes"`
	CpuUsage     float64              `json:"cpu_usage" yaml:"cpu_usage" bson:"cpu_usage"`
	MemUsage     float64              `json:"mem_usage" yaml:"mem_usage" bson:"mem_usage"`
	CpuLimit     float64              `json:"cpu_limit" yaml:"cpu_limit" bson:"cpu_limit"`
	MemLimit     float64              `json:"mem_limit" yaml:"mem_limit" bson:"mem_limit"`
	NfLoad       float64              `json:"mem_load" yaml:"mem_load" bson:"mem_load"`
	NfStatus     NfStatus             `json:"nf_status" yaml:"nf_status" bson:"nf_status"`
	Confidence   float64              `json:"confidence" yaml:"confidence" bson:"confidence"`
}
