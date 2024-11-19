package models

type EventTable struct {
	ID    int64   `gorm:"primaryKey" json:"id"`
	Event EventId `json:"eventId"`
}
