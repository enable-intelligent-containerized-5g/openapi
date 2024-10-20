package models

type PartitioningCriteria string

// List of PartitioningCriteria
const (
	PartitioningCriteria_TAC     PartitioningCriteria = "TAC"
	PartitioningCriteria_SUBPLMN PartitioningCriteria = "SUBPLMN"
	PartitioningCriteria_GEOAREA PartitioningCriteria = "GEOAREA"
	PartitioningCriteria_SNSSAI  PartitioningCriteria = "SNSSAI"
	PartitioningCriteria_DNN     PartitioningCriteria = "DNN"
)
