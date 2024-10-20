package models

// Identifies time and day of the week when the UE is available for communication.
type ScheduledCommunicationTime struct {
	// Identifies the day(s) of the week. If absent, it indicates every day of the week.
	DaysOfWeek []int32 `json:"daysOfWeek,omitempty" yaml:"daysOfWeek" bson:"daysOfWeek,omitempty"`
	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDayStart string `json:"timeOfDayStart,omitempty" yaml:"timeOfDayStart" bson:"timeOfDayStart,omitempty"`
	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDayEnd string `json:"timeOfDayEnd,omitempty" yaml:"timeOfDayEnd" bson:"timeOfDayEnd,omitempty"`
}
