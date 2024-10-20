package models

// Contains the collective behaviour information to be reported to the subscriber.
type CollectiveBehaviourInfo struct {
	ColAttrib []PerUeAttribute `json:"colAttrib" yaml:"colAttrib" bson:"colAttrib,omitempty"`
	// Total number of UEs that fulfil a collective within the area of interest.
	NoOfUes  int32    `json:"noOfUes,omitempty" yaml:"noOfUes" bson:"noOfUes,omitempty"`
	AppIds   []string `json:"appIds,omitempty" yaml:"appIds" bson:"appIds,omitempty"`
	ExtUeIds []string `json:"extUeIds,omitempty" yaml:"extUeIds" bson:"extUeIds,omitempty"`
	UeIds    []string `json:"ueIds,omitempty" yaml:"ueIds" bson:"ueIds,omitempty"`
}
