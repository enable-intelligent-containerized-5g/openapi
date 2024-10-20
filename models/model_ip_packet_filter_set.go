package models

type IpPacketFilterSet struct {
	SrcIp     string `json:"srcIp,omitempty" yaml:"srcIp" bson:"srcIp,omitempty"`
	DstIp     string `json:"dstIp,omitempty" yaml:"dstIp" bson:"dstIp,omitempty"`
	Protocol  int32  `json:"protocol,omitempty" yaml:"protocol" bson:"protocol,omitempty"`
	SrcPort   int32  `json:"srcPort,omitempty" yaml:"srcPort" bson:"srcPort,omitempty"`
	DstPort   int32  `json:"dstPort,omitempty" yaml:"dstPort" bson:"dstPort,omitempty"`
	ToSTc     string `json:"toSTc,omitempty" yaml:"toSTc" bson:"toSTc,omitempty"`
	FlowLabel int32  `json:"flowLabel,omitempty" yaml:"flowLabel" bson:"flowLabel,omitempty"`
	Spi       int32  `json:"spi,omitempty" yaml:"spi" bson:"spi,omitempty"`
	Direction string `json:"direction" yaml:"direction" bson:"direction,omitempty"`
}
