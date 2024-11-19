package models

type NwdafAnalyticsInfoNfLoad struct {
	Accuracy     NwdafMlModelAccuracy  `json:"accuracy,omitempty" yaml:"accuracy,omitempty" bson:"accuracy,omitempty"`
	NfInstanceId string                `json:"nfInstanceId,omitempty" yaml:"nfInstanceId,omitempty" bson:"nfInstanceId,omitempty"`
	NfType       NfType                `json:"nfType,omitempty" yaml:"nfType,omitempty" bson:"nfType,omitempty"`
	Pod          string                `json:"pod,omitempty" yaml:"pod,omitempty" bson:"pod,omitempty"`
	Container    string                `json:"container,omitempty" yaml:"container,omitempty" bson:"container,omitempty"`
	CpuUsage     float64               `json:"cpuUsage" yaml:"cpuUsage" bson:"cpuUsage"`
	MemUsage     float64               `json:"memUsage" yaml:"memUsage" bson:"memUsage"`
	CpuLimit     float64               `json:"cpuLimit" yaml:"cpuLimit" bson:"cpuLimit"`
	MemLimit     float64               `json:"memLimit" yaml:"memLimit" bson:"memLimit"`
	NfLoad       ResourcesNfLoad       `json:"nfLoad" yaml:"nfLoad" bson:"nfLoad"`
	NfStatus     NfStatus              `json:"nfStatus,omitempty" yaml:"nfStatus,omitempty" bson:"nfStatus,omitempty"`
	Confidence   MlModelDataConfidence `json:"confidence,omitempty" yaml:"confidence,omitempty" bson:"confidence,omitempty"`
}
