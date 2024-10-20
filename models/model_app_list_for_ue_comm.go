package models

import (
	"time"
)

// Represents the analytics of the application list used by UE.
type AppListForUeComm struct {
	// String providing an application identifier.
	AppId string `json:"appId" yaml:"appId" bson:"appId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty" yaml:"startTime" bson:"startTime,omitempty"`
	// indicating a time in seconds.
	AppDur int32 `json:"appDur,omitempty" yaml:"appDur" bson:"appDur,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	OccurRatio      int32            `json:"occurRatio,omitempty" yaml:"occurRatio" bson:"occurRatio,omitempty"`
	SpatialValidity *NetworkAreaInfo `json:"spatialValidity,omitempty" yaml:"spatialValidity" bson:"spatialValidity,omitempty"`
}
