package models

// Represents load level information of a given NF instance.
type NfLoadLevelInformation struct {
	NfType NrfNfManagementNfType `json:"nfType,omitempty" yaml:"nfType" bson:"nfType,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfInstanceId string `json:"nfInstanceId,omitempty" yaml:"nfInstanceId" bson:"nfInstanceId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	NfSetId            string                           `json:"nfSetId,omitempty" yaml:"nfSetId" bson:"nfSetId,omitempty"`
	NfStatus           *NwdafEventsSubscriptionNfStatus `json:"nfStatus,omitempty" yaml:"nfStatus" bson:"nfStatus,omitempty"`
	NfCpuUsage         int32                            `json:"nfCpuUsage,omitempty" yaml:"nfCpuUsage" bson:"nfCpuUsage,omitempty"`
	NfMemoryUsage      int32                            `json:"nfMemoryUsage,omitempty" yaml:"nfMemoryUsage" bson:"nfMemoryUsage,omitempty"`
	NfStorageUsage     int32                            `json:"nfStorageUsage,omitempty" yaml:"nfStorageUsage" bson:"nfStorageUsage,omitempty"`
	NfLoadLevelAverage int32                            `json:"nfLoadLevelAverage,omitempty" yaml:"nfLoadLevelAverage" bson:"nfLoadLevelAverage,omitempty"`
	NfLoadLevelpeak    int32                            `json:"nfLoadLevelpeak,omitempty" yaml:"nfLoadLevelpeak" bson:"nfLoadLevelpeak,omitempty"`
	NfLoadAvgInAoi     int32                            `json:"nfLoadAvgInAoi,omitempty" yaml:"nfLoadAvgInAoi" bson:"nfLoadAvgInAoi,omitempty"`
	Snssai             *Snssai                          `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
}