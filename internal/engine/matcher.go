package engine

import (
	"context"

	"cabral-matching/internal/models"
)

type Matcher interface {
	Match(ctx context.Context, req models.MatchRequest) (models.MatchResponse, error)
}
