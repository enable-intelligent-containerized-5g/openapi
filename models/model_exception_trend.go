package models

type ExceptionTrend string

// List of ExceptionTrend
const (
	ExceptionTrend_UP     ExceptionTrend = "UP"
	ExceptionTrend_DOWN   ExceptionTrend = "DOWN"
	ExceptionTrend_UNKNOW ExceptionTrend = "UNKNOW"
	ExceptionTrend_STABLE ExceptionTrend = "STABLE"
)
