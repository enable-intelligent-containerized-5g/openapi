package models

// Represents the S-NSSAI and the optionally associated Network Slice Instance(s).
type NsiIdInfo struct {
	Snssai *Snssai  `json:"snssai" yaml:"snssai" bson:"snssai,omitempty"`
	NsiIds []string `json:"nsiIds,omitempty" yaml:"nsiIds" bson:"nsiIds,omitempty"`
}
