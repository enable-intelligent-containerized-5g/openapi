package models

type EventTable struct {
	ID    int64 `gorm:"primaryKey"`
	Event EventId
}
