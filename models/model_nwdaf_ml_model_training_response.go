package models

type MlModelTrainingResponse struct {
	EventId      EventId               `json:"eventId,omitempty" yaml:"eventId" bson:"eventId" mapstructure:"eventId" db:"eventId" validate:"required"`
	Name         string                `json:"name,omitempty" yaml:"name" bson:"eventId" mapstructure:"name" db:"name" validate:"required"`
	Size         int64                 `json:"size,omitempty" yaml:"size" bson:"size" mapstructure:"size" db:"size" validate:"required"`
	FigureURI    string                `json:"figureUri,omitempty" yaml:"figureUri" bson:"figureUri" mapstructure:"figureUri" db:"figureUri" validate:"required"`
	TargetPeriod int64                 `json:"targetPeriod,omitempty" yaml:"targetPeriod" bson:"targetPeriod" mapstructure:"targetPeriod" db:"targetPeriod" validate:"required"`
	Confidence   MlModelDataConfidence `json:"confidence,omitempty" yaml:"confidence" bson:"confidence" mapstructure:"confidence" db:"confidence" validate:"required"`
	URI          string                `json:"uri,omitempty" yaml:"uri" bson:"uri" mapstructure:"uri" db:"uri" validate:"required"`
	Accuracy     NwdafMlModelAccuracy  `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy" mapstructure:"accuracy" db:"accuracy" validate:"required"`
	NfType       NfType                `json:"nfType,omitempty" yaml:"nfType" bson:"nfType" mapstructure:"nfType" db:"nfType" validate:"required"`
	Figure       string                `json:"figure,omitempty" yaml:"figure" bson:"figure" mapstructure:"figure" db:"figure" validate:"required"`
}
