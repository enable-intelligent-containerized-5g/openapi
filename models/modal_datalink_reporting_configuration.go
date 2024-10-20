package models

type DatalinkReportingConfiguration struct {
	DddTrafficDes []DddTrafficDescriptor `json:"dddTrafficDes,omitempty" yaml:"dddTrafficDes" bson:"dddTrafficDes,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn           string                 `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Slice         *Snssai                `json:"slice,omitempty" yaml:"slice" bson:"slice,omitempty"`
	DddStatusList []DlDataDeliveryStatus `json:"dddStatusList,omitempty" yaml:"dddStatusList" bson:"dddStatusList,omitempty"`
}
