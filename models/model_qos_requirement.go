package models

// Represents the QoS requirements.
type QosRequirement struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.
	Var5qi int32 `json:"5qi,omitempty" yaml:"5qi" bson:"5qi,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GfbrUl string `json:"gfbrUl,omitempty" yaml:"gfbrUl" bson:"gfbrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GfbrDl  string          `json:"gfbrDl,omitempty" yaml:"gfbrDl" bson:"gfbrDl,omitempty"`
	ResType QosResourceType `json:"resType,omitempty" yaml:"resType" bson:"resType,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Pdb int32 `json:"pdb,omitempty" yaml:"pdb" bson:"pdb,omitempty"`
	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	Per string `json:"per,omitempty" yaml:"per" bson:"per,omitempty"`
}
