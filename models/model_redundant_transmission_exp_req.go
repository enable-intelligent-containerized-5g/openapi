package models

// Represents other redundant transmission experience analytics requirements.
type RedundantTransmissionExpReq struct {
	RedTOrderCriter RedTransExpOrderingCriterion `json:"redTOrderCriter,omitempty" yaml:"redTOrderCriter" bson:"redTOrderCriter,omitempty"`
	Order           MatchingDirection            `json:"order,omitempty" yaml:"order" bson:"order,omitempty"`
}
