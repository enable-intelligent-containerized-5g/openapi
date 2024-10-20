package models

// Contains the Media Streaming Consumption information collected for an UE Application via AF.
type MsConsumptionCollection struct {
	MsConsumps []string `json:"msConsumps" yaml:"msConsumps" bson:"msConsumps,omitempty"`
}
