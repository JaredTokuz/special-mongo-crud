package crud

import (
	"context"

	"github.com/jaredtokuz/mongo-crud/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	FindOne(filter interface{}, opts *options.FindOneOptions) (*interface{}, error)
	Find(filter interface{}, opts *options.FindOptions) (*[]interface{}, error)
	FindParameters(id string) (*entities.CrudParameter, error)
}
type repository struct {
	Database *mongo.Database
}

//NewRepo is the single instance repo that is being created.
func NewRepo(database *mongo.Database) Repository {
	return &repository{
		Database: database,
	}
}

// 
func (r *repository) FindOne(filter interface{}, opts *options.FindOneOptions) (*interface{}, error) {
	result := r.Database.Collection("").FindOne(context.Background(), filter, opts)
	var data interface{}
	err := result.Err()
	if err != nil {
		return nil, err
	}
	_ = result.Decode(&data)
	return &data, nil
}

func (r *repository) Find(filter interface{}, opts *options.FindOptions) (*[]interface{}, error) {
	var data []interface{}
	cursor, err := r.Database.Collection("").Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var obj interface{}
		_ = cursor.Decode(&obj)
		data = append(data, obj)
	}
	return &data, nil
}

// Query to get the crud parameters store in data base
func (r *repository) FindParameters(id string) (*entities.CrudParameter, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
			return nil, err
	}
	var data entities.CrudParameter
	if err := r.Database.Collection("crud-parameters").FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}