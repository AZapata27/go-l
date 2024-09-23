package ports

import (
	"context"
	"github.com/jairogloz/go-l/pkg/domain"
)

// TournamentEditionRepository is the interface that have methods to interact with the tournament edition entity in the database.
type TournamentEditionRepository interface {
	Insert(ctx context.Context, tournamentEdition *domain.TournamentEdition) (err error)
}
