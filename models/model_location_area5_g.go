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

 // Represents a user location area when the UE is attached to 5G.
 type LocationArea5G struct {
	 // Identifies a list of geographic area of the user where the UE is located.
	 GeographicAreas []GeographicArea `json:"geographicAreas,omitempty" yaml:"geographicAreas" bson:"geographicAreas,omitempty"`
	 // Identifies a list of civic addresses of the user where the UE is located.
	 CivicAddresses []CivicAddress   `json:"civicAddresses,omitempty" yaml:"civicAddresses" bson:"civicAddresses,omitempty"`
	 NwAreaInfo     *NetworkAreaInfo `json:"nwAreaInfo,omitempty" yaml:"nwAreaInfo" bson:"nwAreaInfo,omitempty"`
 }