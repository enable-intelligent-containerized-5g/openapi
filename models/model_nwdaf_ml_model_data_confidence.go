package models

type MlModelDataConfidence struct {
	R2           float64 `json:"r2,omitempty" yaml:"r2" bson:"r2" mapstructure:"r2" db:"r2" validate:"required"`
	R2Cpu        float64 `json:"r2Cpu,omitempty" yaml:"r2Cpu" bson:"r2Cpu" mapstructure:"r2Cpu" db:"r2Cpu" validate:"required"`
	R2Mem        float64 `json:"r2Mem,omitempty" yaml:"r2Mem" bson:"r2Mem" mapstructure:"r2Mem" db:"r2Mem" validate:"required"`
	R2Troughput  float64 `json:"r2Thrpt,omitempty" yaml:"r2Thrpt" bson:"r2Thrpt" mapstructure:"r2Thrpt" db:"r2Thrpt" validate:"required"`
	MSE          float64 `json:"mse,omitempty" yaml:"mse" bson:"mse" mapstructure:"mse" db:"mse" validate:"required"`
	MSECpu       float64 `json:"mseCpu,omitempty" yaml:"mseCpu" bson:"mseCpu" mapstructure:"mseCpu" db:"mseCpu" validate:"required"`
	MSEMem       float64 `json:"mseMem,omitempty" yaml:"mseMem" bson:"mseMem" mapstructure:"mseMem" db:"mseMem" validate:"required"`
	MSETroughput float64 `json:"mseThrpt,omitempty" yaml:"mseThrpt" bson:"mseThrpt" mapstructure:"mseThrpt" db:"mseThrpt" validate:"required"`
}
