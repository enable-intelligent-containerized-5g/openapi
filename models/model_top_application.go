package models

// Top application that contributes the most to the traffic.
type TopApplication struct {
	// String providing an application identifier.
	AppId           string    `json:"appId,omitempty" yaml:"appId" bson:"appId,omitempty"`
	IpTrafficFilter *FlowInfo `json:"ipTrafficFilter,omitempty" yaml:"ipTrafficFilter" bson:"ipTrafficFilter,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio int32 `json:"ratio,omitempty" yaml:"ratio" bson:"ratio,omitempty"`
}
