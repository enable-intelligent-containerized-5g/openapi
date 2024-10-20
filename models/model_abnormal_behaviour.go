package models

// Represents the abnormal behaviour information.
type AbnormalBehaviour struct {
	Supis []string   `json:"supis,omitempty" yaml:"supis" bson:"supis,omitempty"`
	Excep *Exception `json:"excep" yaml:"excep" bson:"excep,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn    string  `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio int32 `json:"ratio,omitempty" yaml:"ratio" bson:"ratio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence   int32                  `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
	AddtMeasInfo *AdditionalMeasurement `json:"addtMeasInfo,omitempty" yaml:"addtMeasInfo" bson:"addtMeasInfo,omitempty"`
}
