package models

// Indicates the capability supported by the NWDAF
type NwdafCapability struct {
	AnalyticsAggregation          bool `json:"analyticsAggregation,omitempty" yaml:"analyticsAggregation" bson:"analyticsAggregation,omitempty"`
	AnalyticsMetadataProvisioning bool `json:"analyticsMetadataProvisioning,omitempty" yaml:"analyticsMetadataProvisioning" bson:"analyticsMetadataProvisioning,omitempty"`
}
