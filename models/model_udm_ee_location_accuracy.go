package models

type UdmEeLocationAccuracy string

// List of UdmEELocationAccuracy
const (
	UdmEeLocationAccuracy_CELL_LEVEL     UdmEeLocationAccuracy = "CELL_LEVEL"
	UdmEeLocationAccuracy_RAN_NODE_LEVEL UdmEeLocationAccuracy = "RAN_NODE_LEVEL"
	UdmEeLocationAccuracy_TA_LEVEL       UdmEeLocationAccuracy = "TA_LEVEL"
	UdmEeLocationAccuracy_N3_IWF_LEVEL   UdmEeLocationAccuracy = "N3IWF_LEVEL"
	UdmEeLocationAccuracy_UE_IP          UdmEeLocationAccuracy = "UE_IP"
	UdmEeLocationAccuracy_UE_PORT        UdmEeLocationAccuracy = "UE_PORT"
)
