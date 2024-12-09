package models

type MlModelDatasetCommonColumn string

// "namespace", "pod", "container", "timestamp"
const (
	MlModelDatasetCommonColumn_NAMESPACE MlModelDatasetCommonColumn = "namespace"
	MlModelDatasetCommonColumn_POD       MlModelDatasetCommonColumn = "pod"
	MlModelDatasetCommonColumn_CONTAINER MlModelDatasetCommonColumn = "container"
	MlModelDatasetCommonColumn_TIMESTAMP MlModelDatasetCommonColumn = "timestamp"
)
