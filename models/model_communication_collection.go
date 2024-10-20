package models

import (
	"time"
)

// Contains communication information.
type CommunicationCollection struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime" yaml:"startTime" bson:"startTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTime *time.Time `json:"endTime" yaml:"endTime" bson:"endTime,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	UlVol int64 `json:"ulVol" yaml:"ulVol" bson:"ulVol,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DlVol int64 `json:"dlVol" yaml:"dlVol" bson:"dlVol,omitempty"`
}
