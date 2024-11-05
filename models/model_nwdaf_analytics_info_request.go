package models

import "time"

type NwdafAnalyticsInfoRequest struct {
	StartTime     *time.Time           `json:"startTime" yaml:"startTime" bson:"startTime"`
	EndTime       *time.Time           `json:"endTime" yaml:"endTime" bson:"endTime"`
	EventId       EventId              `json:"eventId" yaml:"eventId" bson:"eventId"`
	Accuracy      NwdafMlModelAccuracy `json:"accuracy" yaml:"accuracy" bson:"accuracy"`
	NfInstanceIds []string             `json:"nfInstanceIds" yaml:"nfInstanceIds" bson:"nfInstanceIds"`
	NfTypes       []NfType             `json:"nfTypes" yaml:"nfTypes" bson:"nfTypes"`
}
