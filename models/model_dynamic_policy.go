package models

// A representation of a Dynamic Policy resource.
type DynamicPolicy struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	DynamicPolicyId string `json:"dynamicPolicyId" yaml:"dynamicPolicyId" bson:"dynamicPolicyId,omitempty"`
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	PolicyTemplateId            string                       `json:"policyTemplateId" yaml:"policyTemplateId" bson:"policyTemplateId,omitempty"`
	ServiceDataFlowDescriptions []ServiceDataFlowDescription `json:"serviceDataFlowDescriptions" yaml:"serviceDataFlowDescriptions" bson:"serviceDataFlowDescriptions,omitempty"`
	MediaType                   MediaType                    `json:"mediaType,omitempty" yaml:"mediaType" bson:"mediaType,omitempty"`
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	ProvisioningSessionId string              `json:"provisioningSessionId" yaml:"provisioningSessionId" bson:"provisioningSessionId,omitempty"`
	QosSpecification      *M5QoSSpecification `json:"qosSpecification,omitempty" yaml:"qosSpecification" bson:"qosSpecification,omitempty"`
	EnforcementMethod     string              `json:"enforcementMethod,omitempty" yaml:"enforcementMethod" bson:"enforcementMethod,omitempty"`
	EnforcementBitRate    int32               `json:"enforcementBitRate,omitempty" yaml:"enforcementBitRate" bson:"enforcementBitRate,omitempty"`
}
