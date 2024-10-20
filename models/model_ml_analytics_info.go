package models

// ML Analytics Filter information supported by the Nnwdaf_MLModelProvision service
type MlAnalyticsInfo struct {
	MlAnalyticsIds   []NwdafEvent `json:"mlAnalyticsIds,omitempty" yaml:"mlAnalyticsIds" bson:"mlAnalyticsIds,omitempty"`
	SnssaiList       []Snssai     `json:"snssaiList,omitempty" yaml:"snssaiList" bson:"snssaiList,omitempty"`
	TrackingAreaList []Tai        `json:"trackingAreaList,omitempty" yaml:"trackingAreaList" bson:"trackingAreaList,omitempty"`
}
