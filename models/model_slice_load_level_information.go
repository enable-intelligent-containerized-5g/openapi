package models

// Contains load level information applicable for one or several slices.
type SliceLoadLevelInformation struct {
	// Load level information of the network slice and the optionally associated network slice  instance.
	LoadLevelInformation int32 `json:"loadLevelInformation" yaml:"loadLevelInformation" bson:"loadLevelInformation,omitempty"`
	// Identification(s) of network slice to which the subscription applies.
	Snssais []Snssai `json:"snssais" yaml:"snssais" bson:"snssais,omitempty"`
}
