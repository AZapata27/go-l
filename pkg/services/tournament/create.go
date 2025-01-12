package tournament

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Create creates a new tournament.
func (s *Service) Create(ctx context.Context, tournament *domain.Tournament) (err error) {
	now := time.Now().UTC()
	tournament.CreatedAt = &now

	err = s.Repo.Insert(ctx, tournament)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Println("Duplicate key error")
			appErr := domain.AppError{
				Code: domain.ErrCodeDuplicateKey,
				Msg:  "error creating tournament: duplicate key error",
			}
			return appErr
		}
		log.Println(err.Error())
		return fmt.Errorf("error creating tournament: %w", err)
	}

	return nil
}
