package models

// A representation of a Network Assistance Session resource.
type NetworkAssistanceSession struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	NaSessionId string `json:"naSessionId" yaml:"naSessionId" bson:"naSessionId,omitempty"`
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	ProvisioningSessionId       string                       `json:"provisioningSessionId" yaml:"provisioningSessionId" bson:"provisioningSessionId,omitempty"`
	ServiceDataFlowDescriptions []ServiceDataFlowDescription `json:"serviceDataFlowDescriptions" yaml:"serviceDataFlowDescriptions" bson:"serviceDataFlowDescriptions,omitempty"`
	MediaType                   MediaType                    `json:"mediaType,omitempty" yaml:"mediaType" bson:"mediaType,omitempty"`
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	PolicyTemplateId string              `json:"policyTemplateId,omitempty" yaml:"policyTemplateId" bson:"policyTemplateId,omitempty"`
	RequestedQoS     *M5QoSSpecification `json:"requestedQoS,omitempty" yaml:"requestedQoS" bson:"requestedQoS,omitempty"`
	RecommendedQoS   *M5QoSSpecification `json:"recommendedQoS,omitempty" yaml:"recommendedQoS" bson:"recommendedQoS,omitempty"`
	// Absolute Uniform Resource Locator, conforming with the \"absolute-URI\" production specified in IETF RFC 3986, section 4.3 in which the scheme part is \"http\" or \"https\". Note that the \"query\" suffix is permitted by this production but the \"fragment\" suffix is not.
	NotficationURL string `json:"notficationURL,omitempty" yaml:"notficationURL" bson:"notficationURL,omitempty"`
}
