package models

// Indicates the dispersion class criterion for fixed, camper and/or traveller UE, and/or the  top-heavy UE dispersion class criterion.
type ClassCriterion struct {
	DisperClass *DispersionClass `json:"disperClass" yaml:"disperClass" bson:"disperClass,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	ClassThreshold int32             `json:"classThreshold" yaml:"classThreshold" bson:"classThreshold,omitempty"`
	ThresMatch     MatchingDirection `json:"thresMatch" yaml:"thresMatch" bson:"thresMatch,omitempty"`
}
