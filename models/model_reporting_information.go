package models

import (
	"time"
)

// Represents the type of reporting that the subscription requires.
type ReportingInformation struct {
	ImmRep      bool                               `json:"immRep,omitempty" yaml:"immRep" bson:"immRep,omitempty"`
	NotifMethod SmfEventExposureNotificationMethod `json:"notifMethod,omitempty" yaml:"notifMethod" bson:"notifMethod,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr int32 `json:"maxReportNbr,omitempty" yaml:"maxReportNbr" bson:"maxReportNbr,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	MonDur *time.Time `json:"monDur,omitempty" yaml:"monDur" bson:"monDur,omitempty"`
	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	SampRatio int32 `json:"sampRatio,omitempty" yaml:"sampRatio" bson:"sampRatio,omitempty"`
	// Criteria for partitioning the UEs before applying the sampling ratio.
	PartitionCriteria []PartitioningCriteria `json:"partitionCriteria,omitempty" yaml:"partitionCriteria" bson:"partitionCriteria,omitempty"`
	// indicating a time in seconds.
	GrpRepTime int32            `json:"grpRepTime,omitempty" yaml:"grpRepTime" bson:"grpRepTime,omitempty"`
	NotifFlag  NotificationFlag `json:"notifFlag,omitempty" yaml:"notifFlag" bson:"notifFlag,omitempty"`
}
