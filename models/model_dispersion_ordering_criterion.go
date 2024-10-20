package models

type DispersionOrderingCriterion string

// List of DispersionOrderingCriterion
const (
	DispersionOrderingCriterion_TIME_SLOT_START    DispersionOrderingCriterion = "TIME_SLOT_START"
	DispersionOrderingCriterion_DISPERSION         DispersionOrderingCriterion = "DISPERSION"
	DispersionOrderingCriterion_CLASSIFICATION     DispersionOrderingCriterion = "CLASSIFICATION"
	DispersionOrderingCriterion_RANKING            DispersionOrderingCriterion = "RANKING"
	DispersionOrderingCriterion_PERCENTILE_RANKING DispersionOrderingCriterion = "PERCENTILE_RANKING"
)
