package models

// Represents a QoS flow retainability threshold.
type RetainabilityThreshold struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelFlowNum  int32    `json:"relFlowNum,omitempty" yaml:"relFlowNum" bson:"relFlowNum,omitempty"`
	RelTimeUnit TimeUnit `json:"relTimeUnit,omitempty" yaml:"relTimeUnit" bson:"relTimeUnit,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	RelFlowRatio int32 `json:"relFlowRatio,omitempty" yaml:"relFlowRatio" bson:"relFlowRatio,omitempty"`
}
