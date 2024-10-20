package models

// Contains a list of context identifiers of context information of analytics subscriptions.
type ContextIdList struct {
	ContextIds []AnalyticsContextIdentifier `json:"contextIds" yaml:"contextIds" bson:"contextIds,omitempty"`
}
