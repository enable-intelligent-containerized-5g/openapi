package models

type UdmEeEventType string

// List of UdmEEEventType
const (
	UdmEeEventType_LOSS_OF_CONNECTIVITY                  UdmEeEventType = "LOSS_OF_CONNECTIVITY"
	UdmEeEventType_UE_REACHABILITY_FOR_DATA              UdmEeEventType = "UE_REACHABILITY_FOR_DATA"
	UdmEeEventType_UE_REACHABILITY_FOR_SMS               UdmEeEventType = "UE_REACHABILITY_FOR_SMS"
	UdmEeEventType_LOCATION_REPORTING                    UdmEeEventType = "LOCATION_REPORTING"
	UdmEeEventType_CHANGE_OF_SUPI_PEI_ASSOCIATION        UdmEeEventType = "CHANGE_OF_SUPI_PEI_ASSOCIATION"
	UdmEeEventType_ROAMING_STATUS                        UdmEeEventType = "ROAMING_STATUS"
	UdmEeEventType_COMMUNICATION_FAILURE                 UdmEeEventType = "COMMUNICATION_FAILURE"
	UdmEeEventType_AVAILABILITY_AFTER_DDN_FAILURE        UdmEeEventType = "AVAILABILITY_AFTER_DDN_FAILURE"
	UdmEeEventType_CN_TYPE_CHANGE                        UdmEeEventType = "CN_TYPE_CHANGE"
	UdmEeEventType_DL_DATA_DELIVERY_STATUS               UdmEeEventType = "DL_DATA_DELIVERY_STATUS"
	UdmEeEventType_PDN_CONNECTIVITY_STATUS               UdmEeEventType = "PDN_CONNECTIVITY_STATUS"
	UdmEeEventType_UE_CONNECTION_MANAGEMENT_STATE        UdmEeEventType = "UE_CONNECTION_MANAGEMENT_STATE"
	UdmEeEventType_ACCESS_TYPE_REPORT                    UdmEeEventType = "ACCESS_TYPE_REPORT"
	UdmEeEventType_REGISTRATION_STATE_REPORT             UdmEeEventType = "REGISTRATION_STATE_REPORT"
	UdmEeEventType_CONNECTIVITY_STATE_REPORT             UdmEeEventType = "CONNECTIVITY_STATE_REPORT"
	UdmEeEventType_TYPE_ALLOCATION_CODE_REPORT           UdmEeEventType = "TYPE_ALLOCATION_CODE_REPORT"
	UdmEeEventType_FREQUENT_MOBILITY_REGISTRATION_REPORT UdmEeEventType = "FREQUENT_MOBILITY_REGISTRATION_REPORT"
	UdmEeEventType_PDU_SES_REL                           UdmEeEventType = "PDU_SES_REL"
	UdmEeEventType_PDU_SES_EST                           UdmEeEventType = "PDU_SES_EST"
	UdmEeEventType_UE_MEMORY_AVAILABLE_FOR_SMS           UdmEeEventType = "UE_MEMORY_AVAILABLE_FOR_SMS"
)