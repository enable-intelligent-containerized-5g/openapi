package models

type NwdafMlModelAccuracy string

// List of NwdafEventsSubscriptionAccuracy
const (
	NwdafMlModelAccuracy_LOW    NwdafMlModelAccuracy = "LOW"
	NwdafMlModelAccuracy_MEDIUM NwdafMlModelAccuracy = "MEDIUM"
	NwdafMlModelAccuracy_HIGH   NwdafMlModelAccuracy = "HIGH"
)

// Constructor for asigning default priority to struct
func NewNwdafMlModelAccuracyPriority() []NwdafMlModelAccuracy {
	return []NwdafMlModelAccuracy{
		NwdafMlModelAccuracy_LOW,
		NwdafMlModelAccuracy_MEDIUM,
		NwdafMlModelAccuracy_HIGH,
	}
}
