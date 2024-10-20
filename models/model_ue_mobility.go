package models

import (
	"time"
)

// Represents UE mobility information.
type UeMobility struct {
	// string with format 'date-time' as defined in OpenAPI.
	Ts            *time.Time                  `json:"ts,omitempty" yaml:"ts" bson:"ts,omitempty"`
	RecurringTime *ScheduledCommunicationTime `json:"recurringTime,omitempty" yaml:"recurringTime" bson:"recurringTime,omitempty"`
	// indicating a time in seconds.
	Duration int32 `json:"duration,omitempty" yaml:"duration" bson:"duration,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	DurationVariance float32                               `json:"durationVariance,omitempty" yaml:"durationVariance" bson:"durationVariance,omitempty"`
	LocInfos         []NwdafEventsSubscriptionLocationInfo `json:"locInfos,omitempty" yaml:"locInfos" bson:"locInfos,omitempty"`
}
