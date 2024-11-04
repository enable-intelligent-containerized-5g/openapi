package models

type MlModelInfoData struct {
	URI          string `json:"uri,omitempty" yaml:"uri" bson:"uri" mapstructure:"uri"`
	Accuracy     string `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy" mapstructure:"accuracy"`
	NfType       string `json:"nfType,omitempty" yaml:"nfType" bson:"nfType" mapstructure:"nfType"`
	TargetPeriod string `json:"targetPeriod,omitempty" yaml:"targetPeriod" bson:"targetPeriod" mapstructure:"targetPeriod"`
	EventId      string `json:"eventId,omitempty" yaml:"eventId" bson:"eventId" mapstructure:"eventId"`
}
