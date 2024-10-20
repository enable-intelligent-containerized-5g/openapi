package models

import (
	"time"
)

// Contains UE trajectory information.
type UeTrajectoryInfo struct {
	// string with format 'date-time' as defined in OpenAPI.
	Ts       *time.Time    `json:"ts" yaml:"ts" bson:"ts,omitempty"`
	Location *UserLocation `json:"location" yaml:"location" bson:"location,omitempty"`
}
