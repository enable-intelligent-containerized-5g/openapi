package models

type NwdafFailureCode string

// List of NwdafFailureCode
const (
	NwdafFailureCode_UNAVAILABLE_DATA                     NwdafFailureCode = "UNAVAILABLE_DATA"
	NwdafFailureCode_BOTH_STAT_PRED_NOT_ALLOWED           NwdafFailureCode = "BOTH_STAT_PRED_NOT_ALLOWED"
	NwdafFailureCode_UNSATISFIED_REQUESTED_ANALYTICS_TIME NwdafFailureCode = "UNSATISFIED_REQUESTED_ANALYTICS_TIME"
	NwdafFailureCode_OTHER                                NwdafFailureCode = "OTHER"
)
