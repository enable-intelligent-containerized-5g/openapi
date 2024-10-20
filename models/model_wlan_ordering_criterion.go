package models

type WlanOrderingCriterion string

// List of WlanOrderingCriterion
const (
	WlanOrderingCriterion_TIME_SLOT_START WlanOrderingCriterion = "TIME_SLOT_START"
	WlanOrderingCriterion_NUMBER_OF_UES   WlanOrderingCriterion = "NUMBER_OF_UES"
	WlanOrderingCriterion_RSSI            WlanOrderingCriterion = "RSSI"
	WlanOrderingCriterion_RTT             WlanOrderingCriterion = "RTT"
	WlanOrderingCriterion_TRAFFIC_INFO    WlanOrderingCriterion = "TRAFFIC_INFO"
)
