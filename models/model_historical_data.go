package models

import (
	"time"
)

// Contains historical data related to an analytics subscription.
type HistoricalData struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty" yaml:"startTime" bson:"startTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTime *time.Time `json:"endTime,omitempty" yaml:"endTime" bson:"endTime,omitempty"`
	// Information about subscriptions with the data sources.
	SubsWithSources []SpecificDataSubscription `json:"subsWithSources,omitempty" yaml:"subsWithSources" bson:"subsWithSources,omitempty"`
	// Historical data related to the analytics.
	Data []DataNotification `json:"data" yaml:"data" bson:"data,omitempty"`
}
