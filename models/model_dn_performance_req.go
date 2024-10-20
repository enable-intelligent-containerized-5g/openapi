package models

// Represents other DN performance analytics requirements.
type DnPerformanceReq struct {
	DnPerfOrderCriter DnPerfOrderingCriterion `json:"dnPerfOrderCriter,omitempty" yaml:"dnPerfOrderCriter" bson:"dnPerfOrderCriter,omitempty"`
	Order             MatchingDirection       `json:"order,omitempty" yaml:"order" bson:"order,omitempty"`
	ReportThresholds  []ThresholdLevel        `json:"reportThresholds,omitempty" yaml:"reportThresholds" bson:"reportThresholds,omitempty"`
}
