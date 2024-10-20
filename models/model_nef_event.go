package models

type NefEvent string

// List of NefEvent
const (
	NefEvent_SVC_EXPERIENCE           NefEvent = "SVC_EXPERIENCE"
	NefEvent_UE_MOBILITY              NefEvent = "UE_MOBILITY"
	NefEvent_UE_COMM                  NefEvent = "UE_COMM"
	NefEvent_EXCEPTIONS               NefEvent = "EXCEPTIONS"
	NefEvent_USER_DATA_CONGESTION     NefEvent = "USER_DATA_CONGESTION"
	NefEvent_PERF_DATA                NefEvent = "PERF_DATA"
	NefEvent_DISPERSION               NefEvent = "DISPERSION"
	NefEvent_COLLECTIVE_BEHAVIOUR     NefEvent = "COLLECTIVE_BEHAVIOUR"
	NefEvent_MS_QOE_METRICS           NefEvent = "MS_QOE_METRICS"
	NefEvent_MS_CONSUMPTION           NefEvent = "MS_CONSUMPTION"
	NefEvent_MS_NET_ASSIST_INVOCATION NefEvent = "MS_NET_ASSIST_INVOCATION"
	NefEvent_MS_DYN_POLICY_INVOCATION NefEvent = "MS_DYN_POLICY_INVOCATION"
	NefEvent_MS_ACCESS_ACTIVITY       NefEvent = "MS_ACCESS_ACTIVITY"
)
