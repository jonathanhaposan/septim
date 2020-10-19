package repository

import (
	"github.com/jonathanhaposan/septim/septim-backend/component/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	trxCollection *mongo.Collection
}

func NewRepository(dbClient *db.MongoDB) *Repository {
	trxCollection := dbClient.Client.Database("local-dev").Collection("transaction")

	return &Repository{
		trxCollection: trxCollection,
	}
}
