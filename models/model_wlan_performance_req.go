package models

// Represents other WLAN performance analytics requirements.
type WlanPerformanceReq struct {
	SsIds           []string              `json:"ssIds,omitempty" yaml:"ssIds" bson:"ssIds,omitempty"`
	BssIds          []string              `json:"bssIds,omitempty" yaml:"bssIds" bson:"bssIds,omitempty"`
	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty" yaml:"wlanOrderCriter" bson:"wlanOrderCriter,omitempty"`
	Order           MatchingDirection     `json:"order,omitempty" yaml:"order" bson:"order,omitempty"`
}
