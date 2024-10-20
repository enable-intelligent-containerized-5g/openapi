package models

import (
	"time"
)

// Represents the type of reporting that the subscription requires.
type EventReportingRequirement struct {
	Accuracy NwdafEventsSubscriptionAccuracy `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy,omitempty"`
	// Each element indicates the preferred accuracy level per analytics subset. It may be present if the \"listOfAnaSubsets\" attribute is present in the subscription request when the subscription event is NF_LOAD, UE_COMM, DISPERSION, NETWORK_PERFORMANCE, WLAN_PERFORMANCE, DN_PERFORMANCE or SERVICE_EXPERIENCE.
	AccPerSubset []NwdafEventsSubscriptionAccuracy `json:"accPerSubset,omitempty" yaml:"accPerSubset" bson:"accPerSubset,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StartTs *time.Time `json:"startTs,omitempty" yaml:"startTs" bson:"startTs,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTs *time.Time `json:"endTs,omitempty" yaml:"endTs" bson:"endTs,omitempty"`
	// Offset period in units of seconds to the reporting time, if the value is negative means  statistics in the past offset period, otherwise a positive value means prediction in the  future offset period. May be present if the \"repPeriod\" attribute is included within the  \"evtReq\" attribute.
	OffsetPeriod int32 `json:"offsetPeriod,omitempty" yaml:"offsetPeriod" bson:"offsetPeriod,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	SampRatio int32 `json:"sampRatio,omitempty" yaml:"sampRatio" bson:"sampRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxObjectNbr int32 `json:"maxObjectNbr,omitempty" yaml:"maxObjectNbr" bson:"maxObjectNbr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxSupiNbr int32 `json:"maxSupiNbr,omitempty" yaml:"maxSupiNbr" bson:"maxSupiNbr,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeAnaNeeded     *time.Time                   `json:"timeAnaNeeded,omitempty" yaml:"timeAnaNeeded" bson:"timeAnaNeeded,omitempty"`
	AnaMeta           []AnalyticsMetadata          `json:"anaMeta,omitempty" yaml:"anaMeta" bson:"anaMeta,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty" yaml:"anaMetaInd" bson:"anaMetaInd,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty" yaml:"histAnaTimePeriod" bson:"histAnaTimePeriod,omitempty"`
}
