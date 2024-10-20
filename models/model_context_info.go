package models

type ContextInfo struct {
	OrigHeaders    []string `json:"origHeaders,omitempty" yaml:"origHeaders" bson:"origHeaders,omitempty"`
	RequestHeaders []string `json:"requestHeaders,omitempty" yaml:"requestHeaders" bson:"requestHeaders,omitempty"`
}
