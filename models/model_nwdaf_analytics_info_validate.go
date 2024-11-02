package models

var ValidEventIds = map[EventId]bool{
	EventId_LOAD_LEVEL_INFORMATION: true,
	EventId_NETWORK_PERFORMANCE:    true,
	EventId_NF_LOAD:                true,
	EventId_SERVICE_EXPERIENCE:     true,
	EventId_UE_MOBILITY:            true,
	EventId_UE_COMMUNICATION:       true,
	EventId_QOS_SUSTAINABILITY:     true,
	EventId_ABNORMAL_BEHAVIOUR:     true,
	EventId_USER_DATA_CONGESTION:   true,
	EventId_NSI_LOAD_LEVEL:         true,
	EventId_DN_PERFORMANCE:         true,
	EventId_DISPERSION:             true,
	EventId_RED_TRANS_EXP:          true,
	EventId_WLAN_PERFORMANCE:       true,
	EventId_SM_CONGESTION:          true,
}
