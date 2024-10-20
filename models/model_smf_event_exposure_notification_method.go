package models

type SmfEventExposureNotificationMethod string

// List of SmfEventExposureNotificationMethod
const (
	SmfEventExposureNotificationMethod_PERIODIC           SmfEventExposureNotificationMethod = "PERIODIC"
	SmfEventExposureNotificationMethod_ONE_TIME           SmfEventExposureNotificationMethod = "ONE_TIME"
	SmfEventExposureNotificationMethod_ON_EVENT_DETECTION SmfEventExposureNotificationMethod = "ON_EVENT_DETECTION"
)
