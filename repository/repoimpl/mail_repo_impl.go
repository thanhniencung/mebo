package repoimpl

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mibo/model"

	repo "mibo/repository"
)

// ref >> https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

type MailRepoImpl struct {
	Db *mongo.Database
}

func NewUserRepo(db *mongo.Database) repo.MailRepo {
	return &MailRepoImpl{
		Db: db,
	}
}

// Save email when it was sent
func (m *MailRepoImpl) Save(email model.History) error {
	bbytes, _ := bson.Marshal(email)

	_, err := m.Db.Collection("emails").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}

	return nil
}

// List all email sent
func (m *MailRepoImpl) List() ([]model.History, error) {
	var results []model.History

	findOptions := options.Find()
	findOptions.SetLimit(100)

	// select all
	cur, err := m.Db.Collection("emails").Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		log.Error(err)
		return results, err
	}

	for cur.Next(context.Background()) {
		var elem model.History
		err := cur.Decode(&elem)
		if err != nil {
			break
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Error(err)
		return results, err
	}

	cur.Close(context.Background())

	return results, nil
}
