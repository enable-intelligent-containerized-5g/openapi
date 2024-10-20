package models

// Represents the congestion information.
type CongestionInfo struct {
	CongType  NwdafEventsSubscriptionCongestionType `json:"congType" yaml:"congType" bson:"congType,omitempty"`
	TimeIntev *TimeWindow                           `json:"timeIntev" yaml:"timeIntev" bson:"timeIntev,omitempty"`
	Nsi       *ThresholdLevel                       `json:"nsi" yaml:"nsi" bson:"nsi,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence   int32            `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
	TopAppListUl []TopApplication `json:"topAppListUl,omitempty" yaml:"topAppListUl" bson:"topAppListUl,omitempty"`
	TopAppListDl []TopApplication `json:"topAppListDl,omitempty" yaml:"topAppListDl" bson:"topAppListDl,omitempty"`
}
