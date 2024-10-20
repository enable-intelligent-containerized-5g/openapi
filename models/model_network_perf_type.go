package models

type NetworkPerfType string

// List of NetworkPerfType
const (
	NetworkPerfType_GNB_ACTIVE_RATIO    NetworkPerfType = "GNB_ACTIVE_RATIO"
	NetworkPerfType_GNB_COMPUTING_USAGE NetworkPerfType = "GNB_COMPUTING_USAGE"
	NetworkPerfType_GNB_MEMORY_USAGE    NetworkPerfType = "GNB_MEMORY_USAGE"
	NetworkPerfType_GNB_DISK_USAGE      NetworkPerfType = "GNB_DISK_USAGE"
	NetworkPerfType_NUM_OF_UE           NetworkPerfType = "NUM_OF_UE"
	NetworkPerfType_SESS_SUCC_RATIO     NetworkPerfType = "SESS_SUCC_RATIO"
	NetworkPerfType_HO_SUCC_RATIO       NetworkPerfType = "HO_SUCC_RATIO"
)
