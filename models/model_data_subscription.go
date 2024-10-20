package models

// Contains a data specification.
type DataSubscription struct {
	AmfDataSub   *AmfEventSubscription            `json:"amfDataSub,omitempty" yaml:"amfDataSub" bson:"amfDataSub,omitempty"`
	SmfDataSub   *NsmfEventExposure               `json:"smfDataSub,omitempty" yaml:"smfDataSub" bson:"smfDataSub,omitempty"`
	UdmDataSub   *UdmEeEeSubscription             `json:"udmDataSub,omitempty" yaml:"udmDataSub" bson:"udmDataSub,omitempty"`
	AfDataSub    *AfEventExposureSubsc            `json:"afDataSub,omitempty" yaml:"afDataSub" bson:"afDataSub,omitempty"`
	NefDataSub   *NefEventExposureSubsc           `json:"nefDataSub,omitempty" yaml:"nefDataSub" bson:"nefDataSub,omitempty"`
	NrfDataSub   *NrfNfManagementSubscriptionData `json:"nrfDataSub,omitempty" yaml:"nrfDataSub" bson:"nrfDataSub,omitempty"`
	NsacfDataSub *SacEventSubscription            `json:"nsacfDataSub,omitempty" yaml:"nsacfDataSub" bson:"nsacfDataSub,omitempty"`
}
