package models

type AccuracyTable struct {
	ID       int64 `gorm:"primaryKey"`
	Accuracy NwdafMlModelAccuracy
}
