package models

type LossOfConnectivityReason string

// List of LossOfConnectivityReason
const (
	LossOfConnectivityReason_DEREGISTERED               LossOfConnectivityReason = "DEREGISTERED"
	LossOfConnectivityReason_MAX_DETECTION_TIME_EXPIRED LossOfConnectivityReason = "MAX_DETECTION_TIME_EXPIRED"
	LossOfConnectivityReason_PURGED                     LossOfConnectivityReason = "PURGED"
)
