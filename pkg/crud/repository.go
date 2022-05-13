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
	FindOne(collectionName string, filter interface{}, opts *options.FindOneOptions) (*primitive.M, error)
	Find(collectionName string, filter interface{}, opts *options.FindOptions) (*[]primitive.M, error)
	FindParameters(id string) (*entities.CrudParameter, error)
	DeleteOne(collectionName string, filter interface{}, opts *options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(collectionName string, filter interface{}, opts *options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateOne(collectionName string, filter interface{}, update interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(collectionName string, filter interface{}, update interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error)
	InsertOne(collectionName string, document interface{}, opts *options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(collectionName string, document []interface{}, opts *options.InsertManyOptions) (*mongo.InsertManyResult, error)
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
func (r *repository) FindOne(collectionName string, filter interface{}, opts *options.FindOneOptions) (*primitive.M, error) {
	result := r.Database.Collection(collectionName).FindOne(context.TODO(), filter, opts)
	var data primitive.M
	err := result.Err()
	if err != nil {
		return nil, err
	}
	_ = result.Decode(&data)
	return &data, nil
}

func (r *repository) Find(collectionName string, filter interface{}, opts *options.FindOptions) (*[]primitive.M, error) {
	var data []primitive.M
	cursor, err := r.Database.Collection(collectionName).Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var obj primitive.M
		_ = cursor.Decode(&obj)
		data = append(data, obj)
	}
	return &data, nil
}

func (r *repository) InsertOne(collectionName string, document interface{}, opts *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	result, err := r.Database.Collection(collectionName).InsertOne(context.TODO(), document, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) InsertMany(collectionName string, document []interface{}, opts *options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	result, err := r.Database.Collection(collectionName).InsertMany(context.TODO(), document, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateOne(collectionName string, filter interface{}, update interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := r.Database.Collection(collectionName).UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateMany(collectionName string, filter interface{}, update interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := r.Database.Collection(collectionName).UpdateMany(context.TODO(), filter, update, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) DeleteOne(collectionName string, filter interface{}, opts *options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := r.Database.Collection(collectionName).DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *repository) DeleteMany(collectionName string, filter interface{}, opts *options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := r.Database.Collection(collectionName).DeleteMany(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	return result, err
}


// Query to get the crud parameters store in data base
func (r *repository) FindParameters(id string) (*entities.CrudParameter, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
			return nil, err
	}
	var data entities.CrudParameter
	if err := r.Database.Collection("crud-parameters").FindOne(
		context.TODO(), 
		bson.M{"_id": objectId}).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}