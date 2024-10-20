package models

type M5QoSSpecification struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDlBitRate string `json:"marBwDlBitRate" yaml:"marBwDlBitRate" bson:"marBwDlBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUlBitRate string `json:"marBwUlBitRate" yaml:"marBwUlBitRate" bson:"marBwUlBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwDlBitRate string `json:"minDesBwDlBitRate,omitempty" yaml:"minDesBwDlBitRate" bson:"minDesBwDlBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwUlBitRate string `json:"minDesBwUlBitRate,omitempty" yaml:"minDesBwUlBitRate" bson:"minDesBwUlBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwDlBitRate string `json:"mirBwDlBitRate" yaml:"mirBwDlBitRate" bson:"mirBwDlBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwUlBitRate string `json:"mirBwUlBitRate" yaml:"mirBwUlBitRate" bson:"mirBwUlBitRate,omitempty"`
	DesLatency     int32  `json:"desLatency,omitempty" yaml:"desLatency" bson:"desLatency,omitempty"`
	DesLoss        int32  `json:"desLoss,omitempty" yaml:"desLoss" bson:"desLoss,omitempty"`
}
