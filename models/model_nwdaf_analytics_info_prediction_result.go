package models

type PredictionResult struct {
	CpuAverage float64 `json:"cpuAverage" yaml:"cpuAverage" bson:"cpuAverage" `
	MemAverage float64 `json:"memAverage" yaml:"memAverage" bson:"memAverage" `
}