package engine

import "cabral-matching/internal/models"

type Ranker interface {
	Rank(results []models.MatchResult, limit int) []models.MatchResult
}
