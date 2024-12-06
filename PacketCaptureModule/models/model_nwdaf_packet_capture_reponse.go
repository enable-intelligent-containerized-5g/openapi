package models

import "github.com/enable-intelligent-containerized-5g/openapi/models"

type NwdafPacketCaptureResponse struct {
	Metrics   []NwdafPacketCaptureMetric `json:"metrics" yaml:"metrics" bson:"metrics"`
	NfType    models.NfType              `json:"nfType,omitempty" yaml:"nfType" bson:"nfType" mapstructure:"nfType"`
	NfService string                     `json:"nfService,omitempty" yaml:"nfService" bson:"nfService" mapstructure:"nfService"`
}
