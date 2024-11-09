package db

import (
	"context"
	"easyjobBackend/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MentorStore interface {
	GetAllMentors(context.Context) (*([]types.Mentor), error)
}

type MongoMentorStore struct {
	Coll *mongo.Collection
}

func NewMongoMentorStore(client *mongo.Client) *MongoMentorStore {
	return &MongoMentorStore{
		Coll: client.Database(DBNAME).Collection(MENTORCOLL),
	}
}

func (p *MongoMentorStore) GetAllMentors(ctx context.Context) (*([]types.Mentor), error) {
	var posts []types.Mentor
	cur, err := p.Coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &posts)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
