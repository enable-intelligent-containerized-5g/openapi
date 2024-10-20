package models

// IP address and/or FQDN.
type AddrFqdn struct {
	IpAddr *IpAddr `json:"ipAddr,omitempty" yaml:"ipAddr" bson:"ipAddr,omitempty"`
	// Indicates an FQDN.
	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn" bson:"fqdn,omitempty"`
}
