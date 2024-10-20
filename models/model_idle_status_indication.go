package models

import (
	"time"
)

// Represents the idle status indication.
type IdleStatusIndication struct {
	// string with format \"date-time\" as defined in OpenAPI.
	TimeStamp *time.Time `json:"timeStamp,omitempty" yaml:"timeStamp" bson:"timeStamp,omitempty"`
	// indicating a time in seconds.
	ActiveTime int32 `json:"activeTime,omitempty" yaml:"activeTime" bson:"activeTime,omitempty"`
	// indicating a time in seconds.
	SubsRegTimer            int32 `json:"subsRegTimer,omitempty" yaml:"subsRegTimer" bson:"subsRegTimer,omitempty"`
	EdrxCycleLength         int32 `json:"edrxCycleLength,omitempty" yaml:"edrxCycleLength" bson:"edrxCycleLength,omitempty"`
	SuggestedNumOfDlPackets int32 `json:"suggestedNumOfDlPackets,omitempty" yaml:"suggestedNumOfDlPackets" bson:"suggestedNumOfDlPackets,omitempty"`
}
