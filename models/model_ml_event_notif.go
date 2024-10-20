package models

// Represents a notification related to a single event that occurred.
type MlEventNotif struct {
	Event           NwdafEvent       `json:"event" yaml:"event" bson:"event,omitempty"`
	NotifCorreId    string           `json:"notifCorreId,omitempty" yaml:"notifCorreId" bson:"notifCorreId,omitempty"`
	MLFileAddr      *MlModelAddr     `json:"mLFileAddr" yaml:"mLFileAddr" bson:"mLFileAddr,omitempty"`
	ValidityPeriod  *TimeWindow      `json:"validityPeriod,omitempty" yaml:"validityPeriod" bson:"validityPeriod,omitempty"`
	SpatialValidity *NetworkAreaInfo `json:"spatialValidity,omitempty" yaml:"spatialValidity" bson:"spatialValidity,omitempty"`
}
