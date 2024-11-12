package models

type ResourcesNfLoad struct {
	CpuLoad float64 `json:"cpuLoad" yaml:"cpuLoad" bson:"cpuLoad"`
	MemLoad float64 `json:"memLoad" yaml:"memLoad" bson:"memLoad"`
}
