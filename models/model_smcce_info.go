package models

// Represents the Session Management congestion control experience information.
type SmcceInfo struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn         string       `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Snssai      *Snssai      `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	SmcceUeList *SmcceUeList `json:"smcceUeList" yaml:"smcceUeList" bson:"smcceUeList,omitempty"`
}
