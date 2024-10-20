package models

import (
	"time"
)

// Represents a notification on events that occurred.
type NwdafEventsSubscriptionEventNotification struct {
	Event NwdafEvent `json:"event" yaml:"event" bson:"event,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Start *time.Time `json:"start,omitempty" yaml:"start" bson:"start,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty" yaml:"expiry" bson:"expiry,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStampGen   *time.Time       `json:"timeStampGen,omitempty" yaml:"timeStampGen" bson:"timeStampGen,omitempty"`
	FailNotifyCode NwdafFailureCode `json:"failNotifyCode,omitempty" yaml:"failNotifyCode" bson:"failNotifyCode,omitempty"`
	// indicating a time in seconds.
	RvWaitTime         int32                                          `json:"rvWaitTime,omitempty" yaml:"rvWaitTime" bson:"rvWaitTime,omitempty"`
	AnaMetaInfo        *AnalyticsMetadataInfo                         `json:"anaMetaInfo,omitempty" yaml:"anaMetaInfo" bson:"anaMetaInfo,omitempty"`
	NfLoadLevelInfos   []NfLoadLevelInformation                       `json:"nfLoadLevelInfos,omitempty" yaml:"nfLoadLevelInfos" bson:"nfLoadLevelInfos,omitempty"`
	NsiLoadLevelInfos  []NsiLoadLevelInfo                             `json:"nsiLoadLevelInfos,omitempty" yaml:"nsiLoadLevelInfos" bson:"nsiLoadLevelInfos,omitempty"`
	SliceLoadLevelInfo *SliceLoadLevelInformation                     `json:"sliceLoadLevelInfo,omitempty" yaml:"sliceLoadLevelInfo" bson:"sliceLoadLevelInfo,omitempty"`
	SvcExps            []NwdafEventsSubscriptionServiceExperienceInfo `json:"svcExps,omitempty" yaml:"svcExps" bson:"svcExps,omitempty"`
	QosSustainInfos    []QosSustainabilityInfo                        `json:"qosSustainInfos,omitempty" yaml:"qosSustainInfos" bson:"qosSustainInfos,omitempty"`
	UeComms            []UeCommunication                              `json:"ueComms,omitempty" yaml:"ueComms" bson:"ueComms,omitempty"`
	UeMobs             []UeMobility                                   `json:"ueMobs,omitempty" yaml:"ueMobs" bson:"ueMobs,omitempty"`
	UserDataCongInfos  []UserDataCongestionInfo                       `json:"userDataCongInfos,omitempty" yaml:"userDataCongInfos" bson:"userDataCongInfos,omitempty"`
	AbnorBehavrs       []AbnormalBehaviour                            `json:"abnorBehavrs,omitempty" yaml:"abnorBehavrs" bson:"abnorBehavrs,omitempty"`
	NwPerfs            []NetworkPerfInfo                              `json:"nwPerfs,omitempty" yaml:"nwPerfs" bson:"nwPerfs,omitempty"`
	DnPerfInfos        []DnPerfInfo                                   `json:"dnPerfInfos,omitempty" yaml:"dnPerfInfos" bson:"dnPerfInfos,omitempty"`
	DisperInfos        []DispersionInfo                               `json:"disperInfos,omitempty" yaml:"disperInfos" bson:"disperInfos,omitempty"`
	RedTransInfos      []RedundantTransmissionExpInfo                 `json:"redTransInfos,omitempty" yaml:"redTransInfos" bson:"redTransInfos,omitempty"`
	WlanInfos          []WlanPerformanceInfo                          `json:"wlanInfos,omitempty" yaml:"wlanInfos" bson:"wlanInfos,omitempty"`
	SmccExps           []SmcceInfo                                    `json:"smccExps,omitempty" yaml:"smccExps" bson:"smccExps,omitempty"`
}
