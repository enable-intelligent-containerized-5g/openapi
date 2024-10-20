package models

import (
	"time"
)

// Contains Performance Data Analytics related information collection.
type PerformanceDataInfo struct {
	// String providing an application identifier.
	AppId           string           `json:"appId,omitempty" yaml:"appId" bson:"appId,omitempty"`
	UeIpAddr        *IpAddr          `json:"ueIpAddr,omitempty" yaml:"ueIpAddr" bson:"ueIpAddr,omitempty"`
	IpTrafficFilter *FlowInfo        `json:"ipTrafficFilter,omitempty" yaml:"ipTrafficFilter" bson:"ipTrafficFilter,omitempty"`
	UserLoc         *UserLocation    `json:"userLoc,omitempty" yaml:"userLoc" bson:"userLoc,omitempty"`
	AppLocs         []string         `json:"appLocs,omitempty" yaml:"appLocs" bson:"appLocs,omitempty"`
	AsAddr          *AddrFqdn        `json:"asAddr,omitempty" yaml:"asAddr" bson:"asAddr,omitempty"`
	PerfData        *PerformanceData `json:"perfData" yaml:"perfData" bson:"perfData,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp *time.Time `json:"timeStamp" yaml:"timeStamp" bson:"timeStamp,omitempty"`
}
