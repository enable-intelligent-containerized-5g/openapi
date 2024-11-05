package models

import "time"

type NwdafAnalyticsInfoRequest struct {
	StartTime     *time.Time           `json:"startTime" yaml:"startTime" bson:"startTime" validate:"required"`
	EndTime       *time.Time           `json:"endTime" yaml:"endTime" bson:"endTime" validate:"required"`
	EventId       EventId              `json:"eventId" yaml:"eventId" bson:"eventId" validate:"required"`
	Accuracy      NwdafMlModelAccuracy `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy,omitempty"`
	NfInstanceIds []string             `json:"nfInstanceIds,omitempty" yaml:"nfInstanceIds" bson:"nfInstanceIds,omitempty"`
	NfTypes       []NfType             `json:"nfTypes,omitempty" yaml:"nfTypes" bson:"nfTypes,omitempty"`
}
