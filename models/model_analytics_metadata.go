package models

type AnalyticsMetadata string

// List of AnalyticsMetadata
const (
	AnalyticsMetadata_NUM_OF_SAMPLES  AnalyticsMetadata = "NUM_OF_SAMPLES"
	AnalyticsMetadata_DATA_WINDOW     AnalyticsMetadata = "DATA_WINDOW"
	AnalyticsMetadata_DATA_STAT_PROPS AnalyticsMetadata = "DATA_STAT_PROPS"
	AnalyticsMetadata_STRATEGY        AnalyticsMetadata = "STRATEGY"
	AnalyticsMetadata_ACCURACY        AnalyticsMetadata = "ACCURACY"
)
