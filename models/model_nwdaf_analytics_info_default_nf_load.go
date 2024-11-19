package models

type DefaultNfLoad struct {
	CpuUsage float64 `json:"cpuUsage,omitempty" yaml:"cpuUsage,omitempty" bson:"cpuUsage,omitempty" `
	MemUsage float64 `json:"memUsage,omitempty" yaml:"memUsage,omitempty" bson:"memUsage,omitempty" `
	CpuLimit float64 `json:"cpuLimit,omitempty" yaml:"cpuLimit,omitempty" bson:"cpuLimit,omitempty" `
	MemLimit float64 `json:"memLimit,omitempty" yaml:"memLimit,omitempty" bson:"memLimit,omitempty" `
	NfLoad   float64 `json:"nfLoad,omitempty" yaml:"nfLoad,omitempty" bson:"nfLoad,omitempty" `
}