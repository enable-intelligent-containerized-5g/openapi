package models

type DlDataDeliveryStatus string

// List of DlDataDeliveryStatus
const (
	DlDataDeliveryStatus_BUFFERED    DlDataDeliveryStatus = "BUFFERED"
	DlDataDeliveryStatus_TRANSMITTED DlDataDeliveryStatus = "TRANSMITTED"
	DlDataDeliveryStatus_DISCARDED   DlDataDeliveryStatus = "DISCARDED"
)
