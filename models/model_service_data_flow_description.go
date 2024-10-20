package models

type ServiceDataFlowDescription struct {
	FlowDescription *IpPacketFilterSet `json:"flowDescription,omitempty" yaml:"flowDescription" bson:"flowDescription,omitempty"`
	DomainName      string             `json:"domainName,omitempty" yaml:"domainName" bson:"domainName,omitempty"`
}
