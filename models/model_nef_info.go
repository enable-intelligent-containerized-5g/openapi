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

 // Information of an NEF NF Instance
 type NefInfo struct {
	 // Identity of the NEF
	 NefId                          string                  `json:"nefId,omitempty" yaml:"nefId" bson:"nefId,omitempty"`
	 PfdData                        *NrfNfManagementPfdData `json:"pfdData,omitempty" yaml:"pfdData" bson:"pfdData,omitempty"`
	 AfEeData                       *AfEventExposureData    `json:"afEeData,omitempty" yaml:"afEeData" bson:"afEeData,omitempty"`
	 GpsiRanges                     []IdentityRange         `json:"gpsiRanges,omitempty" yaml:"gpsiRanges" bson:"gpsiRanges,omitempty"`
	 ExternalGroupIdentifiersRanges []IdentityRange         `json:"externalGroupIdentifiersRanges,omitempty" yaml:"externalGroupIdentifiersRanges" bson:"externalGroupIdentifiersRanges,omitempty"`
	 ServedFqdnList                 []string                `json:"servedFqdnList,omitempty" yaml:"servedFqdnList" bson:"servedFqdnList,omitempty"`
	 TaiList                        []Tai                   `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	 TaiRangeList                   []TaiRange              `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
	 DnaiList                       []string                `json:"dnaiList,omitempty" yaml:"dnaiList" bson:"dnaiList,omitempty"`
	 UnTrustAfInfoList              []UnTrustAfInfo         `json:"unTrustAfInfoList,omitempty" yaml:"unTrustAfInfoList" bson:"unTrustAfInfoList,omitempty"`
	 UasNfFunctionalityInd          bool                    `json:"uasNfFunctionalityInd,omitempty" yaml:"uasNfFunctionalityInd" bson:"uasNfFunctionalityInd,omitempty"`
 }