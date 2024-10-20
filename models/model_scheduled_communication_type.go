package models

type ScheduledCommunicationType string

// List of ScheduledCommunicationType
const (
	ScheduledCommunicationType_DOWNLINK_ONLY ScheduledCommunicationType = "DOWNLINK_ONLY"
	ScheduledCommunicationType_UPLINK_ONLY   ScheduledCommunicationType = "UPLINK_ONLY"
	ScheduledCommunicationType_BIDIRECTIONAL ScheduledCommunicationType = "BIDIRECTIONAL"
)
