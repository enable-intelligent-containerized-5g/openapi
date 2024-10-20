package models

type DnPerfOrderingCriterion string

// List of DnPerfOrderingCriterion
const (
	DnPerfOrderingCriterion_AVERAGE_TRAFFIC_RATE     DnPerfOrderingCriterion = "AVERAGE_TRAFFIC_RATE"
	DnPerfOrderingCriterion_MAXIMUM_TRAFFIC_RATE     DnPerfOrderingCriterion = "MAXIMUM_TRAFFIC_RATE"
	DnPerfOrderingCriterion_AVERAGE_PACKET_DELAY     DnPerfOrderingCriterion = "AVERAGE_PACKET_DELAY"
	DnPerfOrderingCriterion_MAXIMUM_PACKET_DELAY     DnPerfOrderingCriterion = "MAXIMUM_PACKET_DELAY"
	DnPerfOrderingCriterion_AVERAGE_PACKET_LOSS_RATE DnPerfOrderingCriterion = "AVERAGE_PACKET_LOSS_RATE"
)
