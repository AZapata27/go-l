package tournament_edition

import (
	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

// Make sure Repository implements ports.TournamentEditionRepository
// at compile time
var _ ports.TournamentEditionRepository = &Repository{}

// Repository is a struct that represents the repository for the tournament entity.
type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}
