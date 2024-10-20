package models

import (
	"time"
)

// Contains context information corresponding with a specific context identifier.
type ContextElement struct {
	ContextId *AnalyticsContextIdentifier `json:"contextId" yaml:"contextId" bson:"contextId,omitempty"`
	// Output analytics for the analytics subscription which have not yet been sent to the  analytics consumer.
	PendAnalytics []NwdafEventsSubscriptionEventNotification `json:"pendAnalytics,omitempty" yaml:"pendAnalytics" bson:"pendAnalytics,omitempty"`
	// Historical output analytics.
	HistAnalytics []NwdafEventsSubscriptionEventNotification `json:"histAnalytics,omitempty" yaml:"histAnalytics" bson:"histAnalytics,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LastOutputTime *time.Time `json:"lastOutputTime,omitempty" yaml:"lastOutputTime" bson:"lastOutputTime,omitempty"`
	// Information about analytics subscriptions that the NWDAF has with other NWDAFs to perform  aggregation.
	AggrSubs []SpecificAnalyticsSubscription `json:"aggrSubs,omitempty" yaml:"aggrSubs" bson:"aggrSubs,omitempty"`
	// Historical data related to the analytics subscription.
	HistData []HistoricalData `json:"histData,omitempty" yaml:"histData" bson:"histData,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	AdrfId string `json:"adrfId,omitempty" yaml:"adrfId" bson:"adrfId,omitempty"`
	// Type(s) of data stored in the ADRF by the NWDAF.
	AdrfDataTypes []AdrfDataType `json:"adrfDataTypes,omitempty" yaml:"adrfDataTypes" bson:"adrfDataTypes,omitempty"`
	// NWDAF identifiers of NWDAF instances used by the NWDAF service consumer when aggregating  multiple analytics subscriptions.
	AggrNwdafIds []string `json:"aggrNwdafIds,omitempty" yaml:"aggrNwdafIds" bson:"aggrNwdafIds,omitempty"`
	// Contains information identifying the ML model(s) that the consumer NWDAF is currently subscribing for the analytics.
	ModelInfo []ModelInfo `json:"modelInfo,omitempty" yaml:"modelInfo" bson:"modelInfo,omitempty"`
}
