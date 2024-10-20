package models

// Contains the dispersion information collected for an AF.
type AfEventExposureDispersionCollection struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi      string          `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	UeAddr    *IpAddr         `json:"ueAddr,omitempty" yaml:"ueAddr" bson:"ueAddr,omitempty"`
	DataUsage *UsageThreshold `json:"dataUsage" yaml:"dataUsage" bson:"dataUsage,omitempty"`
	// Defines a packet filter of an IP flow.
	FlowDesp string `json:"flowDesp,omitempty" yaml:"flowDesp" bson:"flowDesp,omitempty"`
	// String providing an application identifier.
	AppId string   `json:"appId,omitempty" yaml:"appId" bson:"appId,omitempty"`
	Dnais []string `json:"dnais,omitempty" yaml:"dnais" bson:"dnais,omitempty"`
	// indicating a time in seconds.
	AppDur int32 `json:"appDur,omitempty" yaml:"appDur" bson:"appDur,omitempty"`
}
