package models

type NwdafEventsSubscriptionCongestionType string

// List of NwdafEventsSubscriptionCongestionType
const (
	NwdafEventsSubscriptionCongestionType_USER_PLANE             NwdafEventsSubscriptionCongestionType = "USER_PLANE"
	NwdafEventsSubscriptionCongestionType_CONTROL_PLANE          NwdafEventsSubscriptionCongestionType = "CONTROL_PLANE"
	NwdafEventsSubscriptionCongestionType_USER_AND_CONTROL_PLANE NwdafEventsSubscriptionCongestionType = "USER_AND_CONTROL_PLANE"
)
