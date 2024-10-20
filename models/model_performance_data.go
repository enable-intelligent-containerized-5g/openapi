package models

// Contains Performance Data.
type PerformanceData struct {
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Pdb int32 `json:"pdb,omitempty" yaml:"pdb" bson:"pdb,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.
	Plr int32 `json:"plr,omitempty" yaml:"plr" bson:"plr,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ThrputUl string `json:"thrputUl,omitempty" yaml:"thrputUl" bson:"thrputUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ThrputDl string `json:"thrputDl,omitempty" yaml:"thrputDl" bson:"thrputDl,omitempty"`
}
