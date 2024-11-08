package models

type NwdaAnalyticsInfoNfLoad struct {
	Accuracy     NwdafMlModelAccuracy `json:"accuracy,omitempty" yaml:"accuracy,omitempty" bson:"accuracy,omitempty" `
	NfInstanceId string               `json:"nfInstanceId,omitempty" yaml:"nfInstanceId,omitempty" bson:"nfInstanceId,omitempty" `
	NfType       NfType               `json:"nfType,omitempty" yaml:"nfType,omitempty" bson:"nfType,omitempty" `
	CpuUsage     float64              `json:"cpu_usage,omitempty" yaml:"cpu_usage,omitempty" bson:"cpu_usage,omitempty" `
	MemUsage     float64              `json:"mem_usage,omitempty" yaml:"mem_usage,omitempty" bson:"mem_usage,omitempty" `
	CpuLimit     float64              `json:"cpu_limit,omitempty" yaml:"cpu_limit,omitempty" bson:"cpu_limit,omitempty" `
	MemLimit     float64              `json:"mem_limit,omitempty" yaml:"mem_limit,omitempty" bson:"mem_limit,omitempty" `
	NfLoad       float64              `json:"nf_load,omitempty" yaml:"nf_load,omitempty" bson:"nf_load,omitempty" `
	NfStatus     NfStatus             `json:"nf_status,omitempty" yaml:"nf_status,omitempty" bson:"nf_status,omitempty" `
	Confidence   float64              `json:"confidence,omitempty" yaml:"confidence,omitempty" bson:"confidence,omitempty" `
}
