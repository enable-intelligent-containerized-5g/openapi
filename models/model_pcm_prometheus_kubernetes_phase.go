package models

type KubernetesPhase string

const (
	KubernetesPhase_PENDING   KubernetesPhase = "Pending"
	KubernetesPhase_RUNNING   KubernetesPhase = "Running"
	KubernetesPhase_SUCCEEDED KubernetesPhase = "Succeeded"
	KubernetesPhase_FAILED    KubernetesPhase = "Failed"
	KubernetesPhase_UNKNOWN   KubernetesPhase = "Unknown"
)
