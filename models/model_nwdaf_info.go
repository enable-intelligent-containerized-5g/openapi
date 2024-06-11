/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NwdafInfo struct {
	// GroupId                        string          `json:"groupId,omitempty" yaml:"groupId" bson:"groupId" mapstructure:"GroupId"`
	// SupiRanges                     []SupiRange     `json:"supiRanges,omitempty" yaml:"supiRanges" bson:"supiRanges" mapstructure:"SupiRanges"`
	// GpsiRanges                     []IdentityRange `json:"gpsiRanges,omitempty" yaml:"gpsiRanges" bson:"gpsiRanges" mapstructure:"GpsiRanges"`
	// ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty" yaml:"externalGroupIdentifiersRanges" bson:"externalGroupIdentifiersRanges" mapstructure:"ExternalGroupIdentifiersRanges"`
	SupportedDataSets              []DataSetId     `json:"supportedDataSets,omitempty" yaml:"supportedDataSets" bson:"supportedDataSets" mapstructure:"SupportedDataSets"`
}
