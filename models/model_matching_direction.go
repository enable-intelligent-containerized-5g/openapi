package models

type MatchingDirection string

// List of MatchingDirection
const (
	MatchingDirection_ASCENDING  MatchingDirection = "ASCENDING"
	MatchingDirection_DESCENDING MatchingDirection = "DESCENDING"
	MatchingDirection_CROSSED    MatchingDirection = "CROSSED"
)
