package models

// Contains service experience information associated with a service flow.
type ServiceExperienceInfoPerFlow struct {
	SvcExprc  *SvcExperience `json:"svcExprc,omitempty" yaml:"svcExprc" bson:"svcExprc,omitempty"`
	TimeIntev *TimeWindow    `json:"timeIntev,omitempty" yaml:"timeIntev" bson:"timeIntev,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai             string              `json:"dnai,omitempty" yaml:"dnai" bson:"dnai,omitempty"`
	IpTrafficFilter  *FlowInfo           `json:"ipTrafficFilter,omitempty" yaml:"ipTrafficFilter" bson:"ipTrafficFilter,omitempty"`
	EthTrafficFilter *EthFlowDescription `json:"ethTrafficFilter,omitempty" yaml:"ethTrafficFilter" bson:"ethTrafficFilter,omitempty"`
}
