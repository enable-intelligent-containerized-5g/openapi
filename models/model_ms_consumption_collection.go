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

 // Contains the Media Streaming Consumption information collected for an UE Application via AF.
 type MsConsumptionCollection struct {
	 MsConsumps []string `json:"msConsumps" yaml:"msConsumps" bson:"msConsumps,omitempty"`
 }