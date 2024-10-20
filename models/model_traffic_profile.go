package models

type TrafficProfile string

// List of TrafficProfile
const (
	TrafficProfile_SINGLE_TRANS_UL     TrafficProfile = "SINGLE_TRANS_UL"
	TrafficProfile_SINGLE_TRANS_DL     TrafficProfile = "SINGLE_TRANS_DL"
	TrafficProfile_DUAL_TRANS_UL_FIRST TrafficProfile = "DUAL_TRANS_UL_FIRST"
	TrafficProfile_DUAL_TRANS_DL_FIRST TrafficProfile = "DUAL_TRANS_DL_FIRST"
	TrafficProfile_MULTI_TRANS         TrafficProfile = "MULTI_TRANS"
)
