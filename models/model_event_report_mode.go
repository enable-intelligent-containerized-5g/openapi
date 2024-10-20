package models

type EventReportMode string

// List of EventReportMode
const (
	EventReportMode_PERIODIC           EventReportMode = "PERIODIC"
	EventReportMode_ON_EVENT_DETECTION EventReportMode = "ON_EVENT_DETECTION"
)
