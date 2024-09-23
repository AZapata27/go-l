package tournament_edition

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (r *Repository) Insert(ctx context.Context, tournamentEdition *domain.TournamentEdition) (err error) {

	tournamentEdition.ID = "1"

	_, err = r.Collection.InsertOne(context.Background(), tournamentEdition)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("Duplicate key error")
			return fmt.Errorf("%w: error inserting tournament edition: %s",
				domain.ErrDuplicateKey, err.Error())
		}
		log.Println(err.Error())
		return fmt.Errorf("error inserting tournament edition: %w", err)
	}

	return nil

}
