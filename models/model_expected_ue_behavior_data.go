package models

import (
	"time"
)

type ExpectedUeBehaviourData struct {
	StationaryIndication StationaryIndication `json:"stationaryIndication,omitempty" yaml:"stationaryIndication" bson:"stationaryIndication,omitempty"`
	// indicating a time in seconds.
	CommunicationDurationTime int32 `json:"communicationDurationTime,omitempty" yaml:"communicationDurationTime" bson:"communicationDurationTime,omitempty"`
	// indicating a time in seconds.
	PeriodicTime               int32                       `json:"periodicTime,omitempty" yaml:"periodicTime" bson:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty" yaml:"scheduledCommunicationTime" bson:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty" yaml:"scheduledCommunicationType" bson:"scheduledCommunicationType,omitempty"`
	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G.
	ExpectedUmts      []UdmPpLocationArea `json:"expectedUmts,omitempty" yaml:"expectedUmts" bson:"expectedUmts,omitempty"`
	TrafficProfile    TrafficProfile      `json:"trafficProfile,omitempty" yaml:"trafficProfile" bson:"trafficProfile,omitempty"`
	BatteryIndication *BatteryIndication  `json:"batteryIndication,omitempty" yaml:"batteryIndication" bson:"batteryIndication,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty" yaml:"validityTime" bson:"validityTime,omitempty"`
}
