package models

type UmtTime struct {
	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDay string `json:"timeOfDay" yaml:"timeOfDay" bson:"timeOfDay,omitempty"`
	// integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays  shall be indicated with the next higher numbers. 7 shall indicate Sunday.
	DayOfWeek int32 `json:"dayOfWeek" yaml:"dayOfWeek" bson:"dayOfWeek,omitempty"`
}
