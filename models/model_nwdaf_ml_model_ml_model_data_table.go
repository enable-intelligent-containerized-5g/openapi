package models

type MlModelDataTable struct {
	ID           int64         `gorm:"primaryKey" json:"id"`
	Name         string        `json:"name"`
	URI          string        `json:"uri"`
	FigureURI    string        `json:"figureUri"`
	Size         int64         `json:"size"`
	TargetPeriod int64         `json:"targetPeriod"`
	R2           float64       `json:"r2"`
	MSE          float64       `json:"mse"`
	R2Cpu        float64       `json:"r2Cpu"`
	R2Mem        float64       `json:"r2Mem"`
	MSECpu       float64       `json:"mseCpu"`
	MSEMem       float64       `json:"mseMem"`
	NfTypeID     int64         `gorm:"foreignKey:ID" json:"nfTypeId"`
	AccuracyID   int64         `gorm:"foreignKey:ID" json:"accuracyId"`
	EventID      int64         `gorm:"foreignKey:ID" json:"eventId"`
	NfType       NFTypeTable   `json:"nfType"`   // property name in: models.NwdafMLModelDB_NF_TYPE_KEY
	Accuracy     AccuracyTable `json:"accuracy"` // property name in: models.NwdafMLModelDB_ACCURACY_KEY
	Event        EventTable    `json:"event"`    // property name in: models.NwdafMLModelDB_EVENT_ID_KEY
}
