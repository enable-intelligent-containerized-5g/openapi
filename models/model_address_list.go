package models

// Represents a list of IPv4 and/or IPv6 addresses.
type AddressList struct {
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty" yaml:"ipv4Addrs" bson:"ipv4Addrs,omitempty"`
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty" yaml:"ipv6Addrs" bson:"ipv6Addrs,omitempty"`
}
