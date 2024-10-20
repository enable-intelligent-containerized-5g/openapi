package models

// Contains User Data Congestion Analytics related information collection.
type UserDataCongestionCollection struct {
	// String providing an application identifier.
	AppId           string      `json:"appId,omitempty" yaml:"appId" bson:"appId,omitempty"`
	IpTrafficFilter *FlowInfo   `json:"ipTrafficFilter,omitempty" yaml:"ipTrafficFilter" bson:"ipTrafficFilter,omitempty"`
	TimeInterv      *TimeWindow `json:"timeInterv,omitempty" yaml:"timeInterv" bson:"timeInterv,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ThrputUl string `json:"thrputUl,omitempty" yaml:"thrputUl" bson:"thrputUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ThrputDl string `json:"thrputDl,omitempty" yaml:"thrputDl" bson:"thrputDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ThrputPkUl string `json:"thrputPkUl,omitempty" yaml:"thrputPkUl" bson:"thrputPkUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ThrputPkDl string `json:"thrputPkDl,omitempty" yaml:"thrputPkDl" bson:"thrputPkDl,omitempty"`
}
