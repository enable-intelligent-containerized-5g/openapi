package models

// Contains information on the event for which the subscription is not successful.
type FailureEventInfo struct {
	Event       NwdafEvent       `json:"event" yaml:"event" bson:"event,omitempty"`
	FailureCode NwdafFailureCode `json:"failureCode" yaml:"failureCode" bson:"failureCode,omitempty"`
}
