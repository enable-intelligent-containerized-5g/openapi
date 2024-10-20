package models

// Represents the List of UEs classified based on experience level of Session Management  congestion control.
type SmcceUeList struct {
	HighLevel   []string `json:"highLevel,omitempty" yaml:"highLevel" bson:"highLevel,omitempty"`
	MediumLevel []string `json:"mediumLevel,omitempty" yaml:"mediumLevel" bson:"mediumLevel,omitempty"`
	LowLevel    []string `json:"lowLevel,omitempty" yaml:"lowLevel" bson:"lowLevel,omitempty"`
}
