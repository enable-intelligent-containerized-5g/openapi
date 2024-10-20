package models

import (
	"time"
)

type UdmEeReachabilityForSmsReport struct {
	SmsfAccessType AccessType `json:"smsfAccessType" yaml:"smsfAccessType" bson:"smsfAccessType,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	MaxAvailabilityTime *time.Time `json:"maxAvailabilityTime,omitempty" yaml:"maxAvailabilityTime" bson:"maxAvailabilityTime,omitempty"`
}
