package models

// Represents notifications on events that occurred.
type NwdafMlModelProvNotif struct {
	// Notifications about Individual Events.
	EventNotifs []MlEventNotif `json:"eventNotifs" yaml:"eventNotifs" bson:"eventNotifs,omitempty"`
	// String identifying a subscription to the Nnwdaf_MLModelProvision Service.
	SubscriptionId string `json:"subscriptionId" yaml:"subscriptionId" bson:"subscriptionId,omitempty"`
}
