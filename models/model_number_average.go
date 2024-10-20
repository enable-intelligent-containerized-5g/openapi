package models

// Represents average and variance information.
type NumberAverage struct {
	// string with format 'float' as defined in OpenAPI.
	Number float32 `json:"number" yaml:"number" bson:"number,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	Variance float32 `json:"variance" yaml:"variance" bson:"variance,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	Skewness float32 `json:"skewness,omitempty" yaml:"skewness" bson:"skewness,omitempty"`
}
