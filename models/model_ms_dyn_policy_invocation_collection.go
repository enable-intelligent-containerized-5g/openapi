package models

// Contains the Media Streaming Dynamic Policy invocation collected for an UE Application via AF.
type MsDynPolicyInvocationCollection struct {
	MsDynPlyInvocs []DynamicPolicy `json:"msDynPlyInvocs" yaml:"msDynPlyInvocs" bson:"msDynPlyInvocs,omitempty"`
}
