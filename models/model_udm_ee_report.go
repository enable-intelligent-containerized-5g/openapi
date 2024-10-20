package models

type UdmEeReport struct {
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	NewPei              string                   `json:"newPei" yaml:"newPei" bson:"newPei,omitempty"`
	Roaming             bool                     `json:"roaming" yaml:"roaming" bson:"roaming,omitempty"`
	NewServingPlmn      *PlmnId                  `json:"newServingPlmn" yaml:"newServingPlmn" bson:"newServingPlmn,omitempty"`
	AccessType          AccessType               `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	NewCnType           CnType                   `json:"newCnType" yaml:"newCnType" bson:"newCnType,omitempty"`
	OldCnType           CnType                   `json:"oldCnType,omitempty" yaml:"oldCnType" bson:"oldCnType,omitempty"`
	OldCmInfoList       []CmInfo                 `json:"oldCmInfoList,omitempty" yaml:"oldCmInfoList" bson:"oldCmInfoList,omitempty"`
	NewCmInfoList       []CmInfo                 `json:"newCmInfoList" yaml:"newCmInfoList" bson:"newCmInfoList,omitempty"`
	LossOfConnectReason LossOfConnectivityReason `json:"lossOfConnectReason" yaml:"lossOfConnectReason" bson:"lossOfConnectReason,omitempty"`
	Location            *UserLocation            `json:"location" yaml:"location" bson:"location,omitempty"`
	PdnConnStat         PdnConnectivityStatus    `json:"pdnConnStat" yaml:"pdnConnStat" bson:"pdnConnStat,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSeId int32 `json:"pduSeId,omitempty" yaml:"pduSeId" bson:"pduSeId,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr     string         `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr,omitempty"`
	Ipv6Prefixes []string       `json:"ipv6Prefixes,omitempty" yaml:"ipv6Prefixes" bson:"ipv6Prefixes,omitempty"`
	Ipv6Addrs    []string       `json:"ipv6Addrs,omitempty" yaml:"ipv6Addrs" bson:"ipv6Addrs,omitempty"`
	PduSessType  PduSessionType `json:"pduSessType,omitempty" yaml:"pduSessType" bson:"pduSessType,omitempty"`
}
