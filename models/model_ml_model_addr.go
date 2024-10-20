package models

// Addresses of ML model files.
type MlModelAddr struct {
	// String providing an URI formatted according to RFC 3986.
	MLModelUrl string `json:"mLModelUrl,omitempty" yaml:"mLModelUrl" bson:"mLModelUrl,omitempty"`
	// The FQDN of the ML Model file.
	MlFileFqdn string `json:"mlFileFqdn,omitempty" yaml:"mlFileFqdn" bson:"mlFileFqdn,omitempty"`
}
