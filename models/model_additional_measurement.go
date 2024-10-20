package models

import (
	"time"
)

// Represents additional measurement information.
type AdditionalMeasurement struct {
	UnexpLoc      *NetworkAreaInfo          `json:"unexpLoc,omitempty" yaml:"unexpLoc" bson:"unexpLoc,omitempty"`
	UnexpFlowTeps []IpEthFlowDescription    `json:"unexpFlowTeps,omitempty" yaml:"unexpFlowTeps" bson:"unexpFlowTeps,omitempty"`
	UnexpWakes    []time.Time               `json:"unexpWakes,omitempty" yaml:"unexpWakes" bson:"unexpWakes,omitempty"`
	DdosAttack    *AddressList              `json:"ddosAttack,omitempty" yaml:"ddosAttack" bson:"ddosAttack,omitempty"`
	WrgDest       *AddressList              `json:"wrgDest,omitempty" yaml:"wrgDest" bson:"wrgDest,omitempty"`
	Circums       []CircumstanceDescription `json:"circums,omitempty" yaml:"circums" bson:"circums,omitempty"`
}
