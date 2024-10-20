package models

import (
	"time"
)

// Represents a Data Subscription Notification.
type DataNotification struct {
	// List of notifications of AMF events.
	AmfEventNotifs []AmfEventNotification `json:"amfEventNotifs,omitempty" yaml:"amfEventNotifs" bson:"amfEventNotifs,omitempty"`
	// List of notifications of SMF events.
	SmfEventNotifs []NsmfEventExposureNotification `json:"smfEventNotifs,omitempty" yaml:"smfEventNotifs" bson:"smfEventNotifs,omitempty"`
	// List of notifications of UDM events.
	UdmEventNotifs []UdmEeMonitoringReport `json:"udmEventNotifs,omitempty" yaml:"udmEventNotifs" bson:"udmEventNotifs,omitempty"`
	// List of notifications of NEF events.
	NefEventNotifs []NefEventExposureNotif `json:"nefEventNotifs,omitempty" yaml:"nefEventNotifs" bson:"nefEventNotifs,omitempty"`
	// List of notifications of AF events.
	AfEventNotifs []AfEventExposureNotif `json:"afEventNotifs,omitempty" yaml:"afEventNotifs" bson:"afEventNotifs,omitempty"`
	// List of notifications of NRF events.
	NrfEventNotifs []NrfNfManagementNotificationData `json:"nrfEventNotifs,omitempty" yaml:"nrfEventNotifs" bson:"nrfEventNotifs,omitempty"`
	// List of notifications of NSACF events.
	NsacfEventNotifs []SacEventReport `json:"nsacfEventNotifs,omitempty" yaml:"nsacfEventNotifs" bson:"nsacfEventNotifs,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	TimeStamp *time.Time `json:"timeStamp,omitempty" yaml:"timeStamp" bson:"timeStamp,omitempty"`
}
