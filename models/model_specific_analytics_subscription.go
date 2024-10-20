package models

// Represents an existing subscription for a specific type of analytics to a specific NWDAF.
type SpecificAnalyticsSubscription struct {
	SubscriptionId string `json:"subscriptionId,omitempty" yaml:"subscriptionId" bson:"subscriptionId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	ProducerId string `json:"producerId,omitempty" yaml:"producerId" bson:"producerId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	ProducerSetId string                    `json:"producerSetId,omitempty" yaml:"producerSetId" bson:"producerSetId,omitempty"`
	NwdafEvSub    *NnwdafEventsSubscription `json:"nwdafEvSub,omitempty" yaml:"nwdafEvSub" bson:"nwdafEvSub,omitempty"`
}
