package models

// Contains analytics metadata information required for analytics aggregation.
type AnalyticsMetadataInfo struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumSamples    int32                           `json:"numSamples,omitempty" yaml:"numSamples" bson:"numSamples,omitempty"`
	DataWindow    *TimeWindow                     `json:"dataWindow,omitempty" yaml:"dataWindow" bson:"dataWindow,omitempty"`
	DataStatProps []DatasetStatisticalProperty    `json:"dataStatProps,omitempty" yaml:"dataStatProps" bson:"dataStatProps,omitempty"`
	Strategy      OutputStrategy                  `json:"strategy,omitempty" yaml:"strategy" bson:"strategy,omitempty"`
	Accuracy      NwdafEventsSubscriptionAccuracy `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy,omitempty"`
}
