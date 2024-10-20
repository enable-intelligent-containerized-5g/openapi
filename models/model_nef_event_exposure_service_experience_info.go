package models

// Contains service experience information associated with an application.
type NefEventExposureServiceExperienceInfo struct {
	// String providing an application identifier.
	AppId          string                         `json:"appId,omitempty" yaml:"appId" bson:"appId,omitempty"`
	Supis          []string                       `json:"supis,omitempty" yaml:"supis" bson:"supis,omitempty"`
	SvcExpPerFlows []ServiceExperienceInfoPerFlow `json:"svcExpPerFlows" yaml:"svcExpPerFlows" bson:"svcExpPerFlows,omitempty"`
}
