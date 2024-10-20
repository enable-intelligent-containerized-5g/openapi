/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package models

 import (
	 "time"
 )
 
 // Represents information related to an event to be reported.
 type AfEventExposureAfEventNotification struct {
	 Event AfEventExposureAfEvent `json:"event" yaml:"event" bson:"event,omitempty"`
	 // string with format \"date-time\" as defined in OpenAPI.
	 TimeStamp        *time.Time                            `json:"timeStamp" yaml:"timeStamp" bson:"timeStamp,omitempty"`
	 SvcExprcInfos    []ServiceExperienceInfoPerApp         `json:"svcExprcInfos,omitempty" yaml:"svcExprcInfos" bson:"svcExprcInfos,omitempty"`
	 UeMobilityInfos  []UeMobilityCollection                `json:"ueMobilityInfos,omitempty" yaml:"ueMobilityInfos" bson:"ueMobilityInfos,omitempty"`
	 UeCommInfos      []UeCommunicationCollection           `json:"ueCommInfos,omitempty" yaml:"ueCommInfos" bson:"ueCommInfos,omitempty"`
	 ExcepInfos       []ExceptionInfo                       `json:"excepInfos,omitempty" yaml:"excepInfos" bson:"excepInfos,omitempty"`
	 CongestionInfos  []UserDataCongestionCollection        `json:"congestionInfos,omitempty" yaml:"congestionInfos" bson:"congestionInfos,omitempty"`
	 PerfDataInfos    []PerformanceDataCollection           `json:"perfDataInfos,omitempty" yaml:"perfDataInfos" bson:"perfDataInfos,omitempty"`
	 DispersionInfos  []AfEventExposureDispersionCollection `json:"dispersionInfos,omitempty" yaml:"dispersionInfos" bson:"dispersionInfos,omitempty"`
	 CollBhvrInfs     []CollectiveBehaviourInfo             `json:"collBhvrInfs,omitempty" yaml:"collBhvrInfs" bson:"collBhvrInfs,omitempty"`
	 MsQoeMetrInfos   []MsQoeMetricsCollection              `json:"msQoeMetrInfos,omitempty" yaml:"msQoeMetrInfos" bson:"msQoeMetrInfos,omitempty"`
	 MsConsumpInfos   []MsConsumptionCollection             `json:"msConsumpInfos,omitempty" yaml:"msConsumpInfos" bson:"msConsumpInfos,omitempty"`
	 MsNetAssInvInfos []MsNetAssInvocationCollection        `json:"msNetAssInvInfos,omitempty" yaml:"msNetAssInvInfos" bson:"msNetAssInvInfos,omitempty"`
	 MsDynPlyInvInfos []MsDynPolicyInvocationCollection     `json:"msDynPlyInvInfos,omitempty" yaml:"msDynPlyInvInfos" bson:"msDynPlyInvInfos,omitempty"`
	 MsAccActInfos    []MsAccessActivityCollection          `json:"msAccActInfos,omitempty" yaml:"msAccActInfos" bson:"msAccActInfos,omitempty"`
 }