package models

// Represents the Exception information.
type Exception struct {
	ExcepId    ExceptionId    `json:"excepId" yaml:"excepId" bson:"excepId,omitempty"`
	ExcepLevel int32          `json:"excepLevel,omitempty" yaml:"excepLevel" bson:"excepLevel,omitempty"`
	ExcepTrend ExceptionTrend `json:"excepTrend,omitempty" yaml:"excepTrend" bson:"excepTrend,omitempty"`
}
