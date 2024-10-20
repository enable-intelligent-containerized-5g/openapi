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

 // Contains the collective behaviour information to be reported to the subscriber.
 type CollectiveBehaviourInfo struct {
	 ColAttrib []PerUeAttribute `json:"colAttrib" yaml:"colAttrib" bson:"colAttrib,omitempty"`
	 // Total number of UEs that fulfil a collective within the area of interest.
	 NoOfUes  int32    `json:"noOfUes,omitempty" yaml:"noOfUes" bson:"noOfUes,omitempty"`
	 AppIds   []string `json:"appIds,omitempty" yaml:"appIds" bson:"appIds,omitempty"`
	 ExtUeIds []string `json:"extUeIds,omitempty" yaml:"extUeIds" bson:"extUeIds,omitempty"`
	 UeIds    []string `json:"ueIds,omitempty" yaml:"ueIds" bson:"ueIds,omitempty"`
 }