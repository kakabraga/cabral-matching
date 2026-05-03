package repository

import (
	"context"

	"cabral-matching/internal/models"
)

// CandidateFilter carries the pre-filter parameters to the database layer.
// Using a struct instead of individual parameters keeps the interface stable
// as new filter fields are added.
type CandidateFilter struct {
	Tags      []string
	Location  *models.LocationFilter
	Budget    *models.BudgetFilter
	Available *bool
}

type CandidateFinder interface {
	FindByFilter(ctx context.Context, filter CandidateFilter) ([]models.Candidate, error)
}
