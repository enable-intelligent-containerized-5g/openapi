/*
 * Nnef_EventExposure
 *
 * NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.591 V17.7.0; 5G System; Network Exposure Function Southbound Services; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.591/
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

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