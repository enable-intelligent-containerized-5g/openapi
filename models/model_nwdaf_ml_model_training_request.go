package models

import "time"

type NwdafMlModelTrainingRequest struct {
	TargetPeriod int64          `json:"targetPeriod,omitempty" yaml:"targetPeriod,omitempty" bson:"targetPeriod" validate:"required"`
	EventId      EventId `json:"eventId,omitempty" yaml:"eventId,omitempty" bson:"eventId,omitempty" validate:"required"`
	NfType       NfType  `json:"nfType,omitempty" yaml:"nfType,omitempty" bson:"nfType,omitempty" validate:"required"`
	StartTime    time.Time      `json:"startTime,omitempty" yaml:"startTime,omitempty" bson:"startTime,omitempty" validate:"required"`
	NewDataset   bool           `json:"newDataset" yaml:"newDataset" bson:"newDataset"`
}