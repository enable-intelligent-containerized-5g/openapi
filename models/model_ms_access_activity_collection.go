package models

// Contains Media Streaming access activity collected for an UE Application via AF.
type MsAccessActivityCollection struct {
	MsAccActs []MediaStreamingAccessRecord `json:"msAccActs" yaml:"msAccActs" bson:"msAccActs,omitempty"`
}
