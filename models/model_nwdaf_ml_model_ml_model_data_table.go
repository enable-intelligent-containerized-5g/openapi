package models

type MlModelDataTable struct {
	ID           int64 `gorm:"primaryKey"`
	Name         string
	URI          string
	FigureURI    string
	Size         int64
	TargetPeriod int64
	R2           float64
	MSE          float64
	R2Cpu        float64
	R2Mem        float64
	MSECpu       float64
	MSEMem       float64
	NfTypeID     int64                `gorm:"foreignKey:ID"`
	AccuracyID   int64                `gorm:"foreignKey:ID"`
	EventID      int64                `gorm:"foreignKey:ID"`
	NfType       NFTypeTable   // property name in: models.NwdafMLModelDB_NF_TYPE_KEY
	Accuracy     AccuracyTable // property name in: models.NwdafMLModelDB_ACCURACY_KEY
	Event        EventTable    // property name in: models.NwdafMLModelDB_EVENT_ID_KEY
}
