package models

// Represents the user data congestion information.
type UserDataCongestionInfo struct {
	NetworkArea    *NetworkAreaInfo `json:"networkArea" yaml:"networkArea" bson:"networkArea,omitempty"`
	CongestionInfo *CongestionInfo  `json:"congestionInfo" yaml:"congestionInfo" bson:"congestionInfo,omitempty"`
	Snssai         *Snssai          `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
}
