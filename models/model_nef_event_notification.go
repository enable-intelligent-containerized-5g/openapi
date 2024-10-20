package models

import (
	"time"
)

// Represents information related to an event to be reported.
type NefEventNotification struct {
	Event NefEvent `json:"event" yaml:"event" bson:"event,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp        *time.Time                              `json:"timeStamp" yaml:"timeStamp" bson:"timeStamp,omitempty"`
	SvcExprcInfos    []NefEventExposureServiceExperienceInfo `json:"svcExprcInfos,omitempty" yaml:"svcExprcInfos" bson:"svcExprcInfos,omitempty"`
	UeMobilityInfos  []UeMobilityInfo                        `json:"ueMobilityInfos,omitempty" yaml:"ueMobilityInfos" bson:"ueMobilityInfos,omitempty"`
	UeCommInfos      []UeCommunicationInfo                   `json:"ueCommInfos,omitempty" yaml:"ueCommInfos" bson:"ueCommInfos,omitempty"`
	ExcepInfos       []ExceptionInfo                         `json:"excepInfos,omitempty" yaml:"excepInfos" bson:"excepInfos,omitempty"`
	CongestionInfos  []UserDataCongestionCollection          `json:"congestionInfos,omitempty" yaml:"congestionInfos" bson:"congestionInfos,omitempty"`
	PerfDataInfos    []PerformanceDataInfo                   `json:"perfDataInfos,omitempty" yaml:"perfDataInfos" bson:"perfDataInfos,omitempty"`
	DispersionInfos  []AfEventExposureDispersionCollection   `json:"dispersionInfos,omitempty" yaml:"dispersionInfos" bson:"dispersionInfos,omitempty"`
	CollBhvrInfs     []CollectiveBehaviourInfo               `json:"collBhvrInfs,omitempty" yaml:"collBhvrInfs" bson:"collBhvrInfs,omitempty"`
	MsQoeMetrInfos   []MsQoeMetricsCollection                `json:"msQoeMetrInfos,omitempty" yaml:"msQoeMetrInfos" bson:"msQoeMetrInfos,omitempty"`
	MsConsumpInfos   []MsConsumptionCollection               `json:"msConsumpInfos,omitempty" yaml:"msConsumpInfos" bson:"msConsumpInfos,omitempty"`
	MsNetAssInvInfos []MsNetAssInvocationCollection          `json:"msNetAssInvInfos,omitempty" yaml:"msNetAssInvInfos" bson:"msNetAssInvInfos,omitempty"`
	MsDynPlyInvInfos []MsDynPolicyInvocationCollection       `json:"msDynPlyInvInfos,omitempty" yaml:"msDynPlyInvInfos" bson:"msDynPlyInvInfos,omitempty"`
	MsAccActInfos    []MsAccessActivityCollection            `json:"msAccActInfos,omitempty" yaml:"msAccActInfos" bson:"msAccActInfos,omitempty"`
}
