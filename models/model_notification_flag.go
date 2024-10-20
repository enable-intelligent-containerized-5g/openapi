package models

type NotificationFlag string

// List of NotificationFlag
const (
	NotificationFlag_ACTIVATE   NotificationFlag = "ACTIVATE"
	NotificationFlag_DEACTIVATE NotificationFlag = "DEACTIVATE"
	NotificationFlag_RETRIEVAL  NotificationFlag = "RETRIEVAL"
)
