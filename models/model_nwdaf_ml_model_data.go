package models

type MlModelData struct {
	EventId      EventId              `json:"eventId,omitempty" yaml:"eventId" bson:"eventId" mapstructure:"eventId"`
	Size         int64                `json:"size,omitempty" yaml:"size" bson:"size" mapstructure:"size"`
	TargetPeriod int                  `json:"targetPeriod,omitempty" yaml:"targetPeriod" bson:"targetPeriod" mapstructure:"targetPeriod"`
	URI          string               `json:"uri,omitempty" yaml:"uri" bson:"uri" mapstructure:"uri"`
	Accuracy     NwdafMlModelAccuracy `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy" mapstructure:"accuracy"`
	NfType       NfType               `json:"nfType,omitempty" yaml:"nfType" bson:"nfType" mapstructure:"nfType"`
}
