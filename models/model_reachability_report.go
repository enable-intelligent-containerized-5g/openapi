package models

import (
	"time"
)

type ReachabilityReport struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	AmfInstanceId  string         `json:"amfInstanceId,omitempty" yaml:"amfInstanceId" bson:"amfInstanceId,omitempty"`
	AccessTypeList []AccessType   `json:"accessTypeList,omitempty" yaml:"accessTypeList" bson:"accessTypeList,omitempty"`
	Reachability   UeReachability `json:"reachability,omitempty" yaml:"reachability" bson:"reachability,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	MaxAvailabilityTime  *time.Time            `json:"maxAvailabilityTime,omitempty" yaml:"maxAvailabilityTime" bson:"maxAvailabilityTime,omitempty"`
	IdleStatusIndication *IdleStatusIndication `json:"idleStatusIndication,omitempty" yaml:"idleStatusIndication" bson:"idleStatusIndication,omitempty"`
}
