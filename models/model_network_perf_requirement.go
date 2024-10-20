package models

// Represents a network performance requirement.
type NetworkPerfRequirement struct {
	NwPerfType NetworkPerfType `json:"nwPerfType" yaml:"nwPerfType" bson:"nwPerfType,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	RelativeRatio int32 `json:"relativeRatio,omitempty" yaml:"relativeRatio" bson:"relativeRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AbsoluteNum int32 `json:"absoluteNum,omitempty" yaml:"absoluteNum" bson:"absoluteNum,omitempty"`
}