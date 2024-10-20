package models

// Represents notifications on network exposure event(s) that occurred for an Individual Network Exposure Event Subscription resource.
type NefEventExposureNotif struct {
	NotifId     string                 `json:"notifId" yaml:"notifId" bson:"notifId,omitempty"`
	EventNotifs []NefEventNotification `json:"eventNotifs" yaml:"eventNotifs" bson:"eventNotifs,omitempty"`
}
