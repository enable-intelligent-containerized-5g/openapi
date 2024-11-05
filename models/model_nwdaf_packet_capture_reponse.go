package models

type NwdafPacketCaptureResponse struct {
	Metrics   []NwdafPacketCaptureMetric `json:"metrics" yaml:"metrics" bson:"metrics"`
	NfType    NfType                     `json:"nfType,omitempty" yaml:"nfType" bson:"nfType" mapstructure:"nfType"`
	NfService string                     `json:"nfService,omitempty" yaml:"nfService" bson:"nfService" mapstructure:"nfService"`
}
