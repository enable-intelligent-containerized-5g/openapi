package models

// Represents the dispersion analytics requirements.
type DispersionRequirement struct {
	DisperType      *DispersionType             `json:"disperType" yaml:"disperType" bson:"disperType,omitempty"`
	ClassCriters    []ClassCriterion            `json:"classCriters,omitempty" yaml:"classCriters" bson:"classCriters,omitempty"`
	RankCriters     []RankingCriterion          `json:"rankCriters,omitempty" yaml:"rankCriters" bson:"rankCriters,omitempty"`
	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty" yaml:"dispOrderCriter" bson:"dispOrderCriter,omitempty"`
	Order           MatchingDirection           `json:"order,omitempty" yaml:"order" bson:"order,omitempty"`
}
