package models

type MlModelDataTable struct {
	ID           int64 `gorm:"primaryKey"`
	URI          string
	Size         int64
	TargetPeriod int64
	Confidence   float64
	NfTypeID     int64         `gorm:"foreignKey:ID"`
	AccuracyID   int64         `gorm:"foreignKey:ID"`
	EventID      int64         `gorm:"foreignKey:ID"`
	NfType       NFTypeTable   // property name in: models.NwdafMLModelDB_NF_TYPE_KEY
	Accuracy     AccuracyTable // property name in: models.NwdafMLModelDB_ACCURACY_KEY
	Event        EventTable    // property name in: models.NwdafMLModelDB_EVENT_ID_KEY
}
