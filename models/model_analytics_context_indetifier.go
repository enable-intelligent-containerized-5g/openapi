package models

// Contains information about available analytics contexts.
type AnalyticsContextIdentifier struct {
	// The identifier of a subscription.
	SubscriptionId string `json:"subscriptionId,omitempty" yaml:"subscriptionId" bson:"subscriptionId,omitempty"`
	// List of analytics types for which NF related analytics contexts can be retrieved.
	NfAnaCtxts []NwdafEvent `json:"nfAnaCtxts,omitempty" yaml:"nfAnaCtxts" bson:"nfAnaCtxts,omitempty"`
	// List of objects that indicate for which SUPI and analytics types combinations analytics  context can be retrieved.
	UeAnaCtxts []UeAnalyticsContextDescriptor `json:"ueAnaCtxts,omitempty" yaml:"ueAnaCtxts" bson:"ueAnaCtxts,omitempty"`
}
