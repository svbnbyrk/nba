package adapters

import mongodb "github.com/svbnbyrk/nba/pkg/db"

type TeamRepositoryInterface interface {
}

type TeamRepository struct {
	*mongodb.MongoDB
}

func NewTeamRepository(db *mongodb.MongoDB) *TeamRepository {
	return &TeamRepository{db}
}
