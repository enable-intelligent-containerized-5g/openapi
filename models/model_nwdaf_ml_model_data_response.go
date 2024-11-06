package models

type MlModelDataResponse struct {
	MlModels      []MlModelData              `json:"mlModels,omitempty" yaml:"mlModels" bson:"mlModels" mapstructure:"mlModels"`
}
