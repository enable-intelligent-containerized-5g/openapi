package models

type DatasetFileMlModelTrainingRequest struct {
	Data string `json:"data" yaml:"data" bson:"data"`
	Name string `json:"name" yaml:"name" bson:"name"`
}
