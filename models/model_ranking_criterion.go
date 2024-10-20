package models

// Indicates the usage ranking criterion between the high, medium and low usage UE.
type RankingCriterion struct {
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	HighBase int32 `json:"highBase" yaml:"highBase" bson:"highBase,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	LowBase int32 `json:"lowBase" yaml:"lowBase" bson:"lowBase,omitempty"`
}
