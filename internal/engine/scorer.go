package engine

import "cabral-matching/internal/models"

type Scorer interface {
	Calculate(candidate models.Candidate, req models.MatchRequest) models.ScoreBreakdown
}
