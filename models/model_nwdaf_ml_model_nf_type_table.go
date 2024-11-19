package models

type NFTypeTable struct {
	ID     int64  `gorm:"primaryKey" json:"id"`
	NfType NfType `json:"nfType"`
}
