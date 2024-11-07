package models

type NFTypeTable struct {
	ID     int64 `gorm:"primaryKey"`
	NfType NfType
}
