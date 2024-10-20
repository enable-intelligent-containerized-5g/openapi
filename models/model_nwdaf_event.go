package models

type NwdafEvent string

// List of NwdafEvent
const (
	NwdafEvent_SLICE_LOAD_LEVEL     NwdafEvent = "SLICE_LOAD_LEVEL"
	NwdafEvent_NETWORK_PERFORMANCE  NwdafEvent = "NETWORK_PERFORMANCE"
	NwdafEvent_NF_LOAD              NwdafEvent = "NF_LOAD"
	NwdafEvent_SERVICE_EXPERIENCE   NwdafEvent = "SERVICE_EXPERIENCE"
	NwdafEvent_UE_MOBILITY          NwdafEvent = "UE_MOBILITY"
	NwdafEvent_UE_COMMUNICATION     NwdafEvent = "UE_COMMUNICATION"
	NwdafEvent_QOS_SUSTAINABILITY   NwdafEvent = "QOS_SUSTAINABILITY"
	NwdafEvent_ABNORMAL_BEHAVIOUR   NwdafEvent = "ABNORMAL_BEHAVIOUR"
	NwdafEvent_USER_DATA_CONGESTION NwdafEvent = "USER_DATA_CONGESTION"
	NwdafEvent_NSI_LOAD_LEVEL       NwdafEvent = "NSI_LOAD_LEVEL"
	NwdafEvent_DN_PERFORMANCE       NwdafEvent = "DN_PERFORMANCE"
	NwdafEvent_DISPERSION           NwdafEvent = "DISPERSION"
	NwdafEvent_RED_TRANS_EXP        NwdafEvent = "RED_TRANS_EXP"
	NwdafEvent_WLAN_PERFORMANCE     NwdafEvent = "WLAN_PERFORMANCE"
	NwdafEvent_SM_CONGESTION        NwdafEvent = "SM_CONGESTION"
)
