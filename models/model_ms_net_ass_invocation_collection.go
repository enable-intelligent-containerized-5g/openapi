package models

// Contains the Media Streaming Network Assistance invocation collected for an UE Application  via AF.
type MsNetAssInvocationCollection struct {
	MsNetAssInvocs []NetworkAssistanceSession `json:"msNetAssInvocs" yaml:"msNetAssInvocs" bson:"msNetAssInvocs,omitempty"`
}
