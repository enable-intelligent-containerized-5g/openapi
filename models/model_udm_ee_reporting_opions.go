package models

import (
	"time"
)

type UdmEeReportingOptions struct {
	ReportMode      EventReportMode `json:"reportMode,omitempty" yaml:"reportMode" bson:"reportMode,omitempty"`
	MaxNumOfReports int32           `json:"maxNumOfReports,omitempty" yaml:"maxNumOfReports" bson:"maxNumOfReports,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty" yaml:"expiry" bson:"expiry,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	SamplingRatio int32 `json:"samplingRatio,omitempty" yaml:"samplingRatio" bson:"samplingRatio,omitempty"`
	// indicating a time in seconds.
	GuardTime int32 `json:"guardTime,omitempty" yaml:"guardTime" bson:"guardTime,omitempty"`
	// indicating a time in seconds.
	ReportPeriod int32            `json:"reportPeriod,omitempty" yaml:"reportPeriod" bson:"reportPeriod,omitempty"`
	NotifFlag    NotificationFlag `json:"notifFlag,omitempty" yaml:"notifFlag" bson:"notifFlag,omitempty"`
}
