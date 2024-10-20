package models

// Contains the serving core network operator PLMN ID and, for an SNPN, the NID that together with the PLMN ID identifies the SNPN.
type PlmnIdNid struct {
	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mcc string `json:"mcc" yaml:"mcc" bson:"mcc,omitempty"`
	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc" yaml:"mnc" bson:"mnc,omitempty"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid string `json:"nid,omitempty" yaml:"nid" bson:"nid,omitempty"`
}
