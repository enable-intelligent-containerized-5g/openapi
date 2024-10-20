package models

// Contains context information related to analytics subscriptions corresponding with one or  more context identifiers.
type ContextData struct {
	// List of items that contain context information corresponding with a context identifier.
	ContextElems []ContextElement `json:"contextElems" yaml:"contextElems" bson:"contextElems,omitempty"`
}
