package models

type ExpectedAnalyticsType string

// List of ExpectedAnalyticsType
const (
	ExpectedAnalyticsType_MOBILITY            ExpectedAnalyticsType = "MOBILITY"
	ExpectedAnalyticsType_COMMUN              ExpectedAnalyticsType = "COMMUN"
	ExpectedAnalyticsType_MOBILITY_AND_COMMUN ExpectedAnalyticsType = "MOBILITY_AND_COMMUN"
)
