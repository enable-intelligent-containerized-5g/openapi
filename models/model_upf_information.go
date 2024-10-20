package models

// Represents the ID/address/FQDN of the UPF.
type UpfInformation struct {
	UpfId   string    `json:"upfId,omitempty" yaml:"upfId" bson:"upfId,omitempty"`
	UpfAddr *AddrFqdn `json:"upfAddr,omitempty" yaml:"upfAddr" bson:"upfAddr,omitempty"`
}
