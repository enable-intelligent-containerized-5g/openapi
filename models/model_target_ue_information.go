package models

// Identifies the target UE information.
type TargetUeInformation struct {
	AnyUe       bool     `json:"anyUe,omitempty" yaml:"anyUe" bson:"anyUe,omitempty"`
	Supis       []string `json:"supis,omitempty" yaml:"supis" bson:"supis,omitempty"`
	Gpsis       []string `json:"gpsis,omitempty" yaml:"gpsis" bson:"gpsis,omitempty"`
	IntGroupIds []string `json:"intGroupIds,omitempty" yaml:"intGroupIds" bson:"intGroupIds,omitempty"`
}
