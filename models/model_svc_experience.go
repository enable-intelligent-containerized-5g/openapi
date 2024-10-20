package models

// Contains a mean opinion score with the customized range.
type SvcExperience struct {
	// string with format 'float' as defined in OpenAPI.
	Mos float32 `json:"mos,omitempty" yaml:"mos" bson:"mos,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	UpperRange float32 `json:"upperRange,omitempty" yaml:"upperRange" bson:"upperRange,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	LowerRange float32 `json:"lowerRange,omitempty" yaml:"lowerRange" bson:"lowerRange,omitempty"`
}
