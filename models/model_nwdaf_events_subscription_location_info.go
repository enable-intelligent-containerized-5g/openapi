package models

// Represents UE location information.
type NwdafEventsSubscriptionLocationInfo struct {
	Loc *UserLocation `json:"loc" yaml:"loc" bson:"loc,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio int32 `json:"ratio,omitempty" yaml:"ratio" bson:"ratio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
}
