package models

// Represents the event(s) that the subscription is not successful including the failure  reason(s).
type FailureEventInfoForMlModel struct {
	Event       NwdafEvent                       `json:"event" yaml:"event" bson:"event,omitempty"`
	FailureCode NwdafMlModelProvisionFailureCode `json:"failureCode" yaml:"failureCode" bson:"failureCode,omitempty"`
}
