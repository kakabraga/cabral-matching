package transport

import (
	"context"

	"cabral-matching/internal/models"
)

type Publisher interface {
	Publish(ctx context.Context, queueName string, response models.MatchResponse) error
}
