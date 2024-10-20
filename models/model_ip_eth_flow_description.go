package models

// Contains the description of an Uplink and/or Downlink Ethernet flow.
type IpEthFlowDescription struct {
	// Defines a packet filter of an IP flow.
	IpTrafficFilter  string              `json:"ipTrafficFilter,omitempty" yaml:"ipTrafficFilter" bson:"ipTrafficFilter,omitempty"`
	EthTrafficFilter *EthFlowDescription `json:"ethTrafficFilter,omitempty" yaml:"ethTrafficFilter" bson:"ethTrafficFilter,omitempty"`
}
