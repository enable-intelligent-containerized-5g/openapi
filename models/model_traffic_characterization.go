package models

// Identifies the detailed traffic characterization.
type TrafficCharacterization struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn    string  `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// String providing an application identifier.
	AppId  string                 `json:"appId,omitempty" yaml:"appId" bson:"appId,omitempty"`
	FDescs []IpEthFlowDescription `json:"fDescs,omitempty" yaml:"fDescs" bson:"fDescs,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	UlVol int64 `json:"ulVol,omitempty" yaml:"ulVol" bson:"ulVol,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	UlVolVariance float32 `json:"ulVolVariance,omitempty" yaml:"ulVolVariance" bson:"ulVolVariance,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DlVol int64 `json:"dlVol,omitempty" yaml:"dlVol" bson:"dlVol,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	DlVolVariance float32 `json:"dlVolVariance,omitempty" yaml:"dlVolVariance" bson:"dlVolVariance,omitempty"`
}
