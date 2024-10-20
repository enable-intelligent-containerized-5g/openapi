package models

// Contains the percentage of time spent on various NF states.
type NwdafEventsSubscriptionNfStatus struct {
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	StatusRegistered int32 `json:"statusRegistered,omitempty" yaml:"statusRegistered" bson:"statusRegistered,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	StatusUnregistered int32 `json:"statusUnregistered,omitempty" yaml:"statusUnregistered" bson:"statusUnregistered,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	StatusUndiscoverable int32 `json:"statusUndiscoverable,omitempty" yaml:"statusUndiscoverable" bson:"statusUndiscoverable,omitempty"`
}
