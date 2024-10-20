package models

type UdmEeLocationReportingConfiguration struct {
	CurrentLocation bool                  `json:"currentLocation" yaml:"currentLocation" bson:"currentLocation,omitempty"`
	OneTime         bool                  `json:"oneTime,omitempty" yaml:"oneTime" bson:"oneTime,omitempty"`
	Accuracy        UdmEeLocationAccuracy `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy,omitempty"`
	N3gppAccuracy   UdmEeLocationAccuracy `json:"n3gppAccuracy,omitempty" yaml:"n3gppAccuracy" bson:"n3gppAccuracy,omitempty"`
}
