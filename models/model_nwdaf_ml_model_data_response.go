package models

type MlModelDataResponse struct {
	MlModelDataResp      []MlModelData              `json:"mlModelDataResp,omitempty" yaml:"mlModelDataResp" bson:"mlModelDataResp" mapstructure:"mlModelDataResp"`
}
