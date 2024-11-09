package db

import (
	"context"
	"easyjobBackend/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobStore interface {
	InsertJob(context.Context, *types.Job) error
	// GetJobByUserID(context.Context, *primitive.ObjectID) (*([]types.Job), error)
	// UpdateJobImageLinksByID(context.Context, *primitive.ObjectID, *([]string)) error
	GetJobByFilter(context.Context, interface{}) (*([]types.Job), error)
	GetAllJobs(context.Context) (*([]types.Job), error)
}

type MongoJobStore struct {
	Coll *mongo.Collection
}

func NewMongoJobStore(client *mongo.Client) *MongoJobStore {
	return &MongoJobStore{
		Coll: client.Database(DBNAME).Collection(JOBCOLL),
	}
}

func (p *MongoJobStore) InsertJob(ctx context.Context, post *types.Job) error {
	res, err := p.Coll.InsertOne(ctx, post)
	if err != nil {
		return err
	}
	post.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// func (p *MongoJobStore) GetJobByUserID(ctx context.Context, id *primitive.ObjectID) (*([]types.Job), error) {
// 	var posts []types.Job
// 	cur, err := p.Coll.Find(ctx, bson.D{{Key: "userId", Value: *id}})
// 	if err != nil {
// 		if err.Error() == mongo.ErrNoDocuments.Error() {
// 			return nil, fmt.Errorf("empty")
// 		}
// 		return nil, err
// 	}
// 	err = cur.All(ctx, &posts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &posts, nil

// }

// func (p *MongoJobStore) UpdateJobImageLinksByID(ctx context.Context, id *primitive.ObjectID, entry *([]string)) error {
// 	filter := bson.M{"_id": *id}
// 	update := bson.M{"$set": bson.M{"postImages": *entry}}
// 	res, err := p.Coll.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	if res.ModifiedCount < 1 {
// 		return fmt.Errorf("not updated")
// 	}
// 	return nil

// }

func (p *MongoJobStore) GetJobByFilter(ctx context.Context, filter interface{}) (*([]types.Job), error) {
	var posts []types.Job
	cur, err := p.Coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &posts)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
func (p *MongoJobStore) GetAllJobs(ctx context.Context) (*([]types.Job), error) {
	var posts []types.Job
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
