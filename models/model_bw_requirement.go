package models

// Represents bandwidth requirements.
type BwRequirement struct {
	// String providing an application identifier.
	AppId string `json:"appId" yaml:"appId" bson:"appId,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDl string `json:"marBwDl,omitempty" yaml:"marBwDl" bson:"marBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUl string `json:"marBwUl,omitempty" yaml:"marBwUl" bson:"marBwUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwDl string `json:"mirBwDl,omitempty" yaml:"mirBwDl" bson:"mirBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwUl string `json:"mirBwUl,omitempty" yaml:"mirBwUl" bson:"mirBwUl,omitempty"`
}
