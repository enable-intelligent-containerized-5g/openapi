package models

import (
	"time"
)

// UE application data collected per UE.
type PerUeAttribute struct {
	UeDest *LocationArea5G `json:"ueDest,omitempty" yaml:"ueDest" bson:"ueDest,omitempty"`
	Route  string          `json:"route,omitempty" yaml:"route" bson:"route,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	AvgSpeed string `json:"avgSpeed,omitempty" yaml:"avgSpeed" bson:"avgSpeed,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeOfArrival *time.Time `json:"timeOfArrival,omitempty" yaml:"timeOfArrival" bson:"timeOfArrival,omitempty"`
}
