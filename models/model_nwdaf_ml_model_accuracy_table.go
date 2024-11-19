package models

type AccuracyTable struct {
	ID       int64                `gorm:"primaryKey" json:"id"`
	Accuracy NwdafMlModelAccuracy `json:"accuracy"`
}
