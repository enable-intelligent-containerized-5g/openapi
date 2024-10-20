package models

// Represents a threshold level.
type ThresholdLevel struct {
	CongLevel      int32 `json:"congLevel,omitempty" yaml:"congLevel" bson:"congLevel,omitempty"`
	NfLoadLevel    int32 `json:"nfLoadLevel,omitempty" yaml:"nfLoadLevel" bson:"nfLoadLevel,omitempty"`
	NfCpuUsage     int32 `json:"nfCpuUsage,omitempty" yaml:"nfCpuUsage" bson:"nfCpuUsage,omitempty"`
	NfMemoryUsage  int32 `json:"nfMemoryUsage,omitempty" yaml:"nfMemoryUsage" bson:"nfMemoryUsage,omitempty"`
	NfStorageUsage int32 `json:"nfStorageUsage,omitempty" yaml:"nfStorageUsage" bson:"nfStorageUsage,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	AvgTrafficRate string `json:"avgTrafficRate,omitempty" yaml:"avgTrafficRate" bson:"avgTrafficRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxTrafficRate string `json:"maxTrafficRate,omitempty" yaml:"maxTrafficRate" bson:"maxTrafficRate,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AvgPacketDelay int32 `json:"avgPacketDelay,omitempty" yaml:"avgPacketDelay" bson:"avgPacketDelay,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	MaxPacketDelay int32 `json:"maxPacketDelay,omitempty" yaml:"maxPacketDelay" bson:"maxPacketDelay,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.
	AvgPacketLossRate int32 `json:"avgPacketLossRate,omitempty" yaml:"avgPacketLossRate" bson:"avgPacketLossRate,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	SvcExpLevel float32 `json:"svcExpLevel,omitempty" yaml:"svcExpLevel" bson:"svcExpLevel,omitempty"`
}
