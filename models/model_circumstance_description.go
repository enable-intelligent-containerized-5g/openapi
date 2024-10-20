package models

import (
	"time"
)

// Contains the description of a circumstance.
type CircumstanceDescription struct {
	// string with format 'float' as defined in OpenAPI.
	Freq float32 `json:"freq,omitempty" yaml:"freq" bson:"freq,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Tm      *time.Time       `json:"tm,omitempty" yaml:"tm" bson:"tm,omitempty"`
	LocArea *NetworkAreaInfo `json:"locArea,omitempty" yaml:"locArea" bson:"locArea,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	Vol int64 `json:"vol,omitempty" yaml:"vol" bson:"vol,omitempty"`
}
