package models

// Represents the network performance information.
type NetworkPerfInfo struct {
	NetworkArea *NetworkAreaInfo `json:"networkArea,omitempty" yaml:"networkArea" bson:"networkArea,omitempty"`
	NwPerfType  NetworkPerfType  `json:"nwPerfType,omitempty" yaml:"nwPerfType" bson:"nwPerfType,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	RelativeRatio int32 `json:"relativeRatio,omitempty" yaml:"relativeRatio" bson:"relativeRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AbsoluteNum int32 `json:"absoluteNum,omitempty" yaml:"absoluteNum" bson:"absoluteNum,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
}