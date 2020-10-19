package repository

import (
	"context"
	"log"
	"time"

	"github.com/jonathanhaposan/septim/septim-backend/internal/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) GetAllTransaction() (res []*model.Transaction) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := r.trxCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("failed get all transaction", err)
		return nil
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var temp model.Transaction

		err = cur.Decode(&temp)
		if err != nil {
			log.Println("failed decode result", err)
			continue
		}

		res = append(res, &temp)
	}

	if cur.Err() != nil {
		log.Println("failed to scan", cur.Err())
		return nil
	}

	return
}
