package models

// Contains analytics metadata information requested to be used during analytics generation.
type AnalyticsMetadataIndication struct {
	DataWindow    *TimeWindow                  `json:"dataWindow,omitempty" yaml:"dataWindow" bson:"dataWindow,omitempty"`
	DataStatProps []DatasetStatisticalProperty `json:"dataStatProps,omitempty" yaml:"dataStatProps" bson:"dataStatProps,omitempty"`
	Strategy      OutputStrategy               `json:"strategy,omitempty" yaml:"strategy" bson:"strategy,omitempty"`
	AggrNwdafIds  []string                     `json:"aggrNwdafIds,omitempty" yaml:"aggrNwdafIds" bson:"aggrNwdafIds,omitempty"`
}
