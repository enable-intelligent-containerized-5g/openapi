package models

// Contains a Traffic Descriptor.
type DddTrafficDescriptor struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr,omitempty"`
	Ipv6Addr string `json:"ipv6Addr,omitempty" yaml:"ipv6Addr" bson:"ipv6Addr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber,omitempty" yaml:"portNumber" bson:"portNumber,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	MacAddr string `json:"macAddr,omitempty" yaml:"macAddr" bson:"macAddr,omitempty"`
}
