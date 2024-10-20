package models

// Represents the exceptions information provided by the AF.
type ExceptionInfo struct {
	IpTrafficFilter  *FlowInfo           `json:"ipTrafficFilter,omitempty" yaml:"ipTrafficFilter" bson:"ipTrafficFilter,omitempty"`
	EthTrafficFilter *EthFlowDescription `json:"ethTrafficFilter,omitempty" yaml:"ethTrafficFilter" bson:"ethTrafficFilter,omitempty"`
	Exceps           []Exception         `json:"exceps" yaml:"exceps" bson:"exceps,omitempty"`
}
