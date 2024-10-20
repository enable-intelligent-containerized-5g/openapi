package models

import (
	"time"
)

// Represents a subscription to a single event.
type MlEventSubscription struct {
	MLEvent        NwdafEvent                     `json:"mLEvent" yaml:"mLEvent" bson:"mLEvent,omitempty"`
	MLEventFilter  *NwdafAnalyticsInfoEventFilter `json:"mLEventFilter" yaml:"mLEventFilter" bson:"mLEventFilter,omitempty"`
	TgtUe          *TargetUeInformation           `json:"tgtUe,omitempty" yaml:"tgtUe" bson:"tgtUe,omitempty"`
	MLTargetPeriod *TimeWindow                    `json:"mLTargetPeriod,omitempty" yaml:"mLTargetPeriod" bson:"mLTargetPeriod,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpiryTime *time.Time `json:"expiryTime,omitempty" yaml:"expiryTime" bson:"expiryTime,omitempty"`
}
