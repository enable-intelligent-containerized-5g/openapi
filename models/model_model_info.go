package models

// Contains information about an ML model.
type ModelInfo struct {
	AnalyticsId  NwdafEvent    `json:"analyticsId" yaml:"analyticsId" bson:"analyticsId,omitempty"`
	MlModelInfos []MlModelInfo `json:"mlModelInfos" yaml:"mlModelInfos" bson:"mlModelInfos,omitempty"`
}
