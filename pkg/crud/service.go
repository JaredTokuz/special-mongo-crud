package crud

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Service is an interface from which our api module can access our repository of all our models
type Service interface {
	FindOne(id string) (*primitive.M, error)
	Find(id string) (*[]primitive.M, error)
	InsertOne(id string) (*mongo.InsertOneResult, error)
	Insert(id string) (*mongo.InsertManyResult, error)
	DeleteOne(id string) (*mongo.DeleteResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
	DaysOutExpire(collectionName string, fieldName string, daysOut int) (*mongo.DeleteResult, error)
	DeleteFirstNRows(collectionName string, nRows int64) (*mongo.DeleteResult, error)
	UpdateOne(id string) (*mongo.UpdateResult, error)
	Update(id string) (*mongo.UpdateResult, error)
}

type service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FindOne(id string) (*primitive.M, error) {
	opts := options.FindOneOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.FindOne("",filter.Filter, &opts)
}

func (s *service) Find(id string) (*[]primitive.M, error) {
	opts := options.FindOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.Find("",filter.Filter, &opts)
}


func (s *service) DeleteOne(id string) (*mongo.DeleteResult, error) {
	opts := options.DeleteOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.DeleteOne("",filter.Filter, &opts)
}

func (s *service) Delete(id string) (*mongo.DeleteResult, error) {
	opts := options.DeleteOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.DeleteMany("",filter.Filter, &opts)
}

func (s *service) DaysOutExpire(collectionName string, fieldName string, daysOut int) (*mongo.DeleteResult, error) {
	opts := options.DeleteOptions{}
	calc_date := time.Now().AddDate(0, 0, -daysOut)
	return s.repository.DeleteMany(collectionName, bson.M{fieldName: bson.M{"$lt": calc_date } }, &opts)
}

func (s *service) DeleteFirstNRows(collectionName string, nRows int64) (*mongo.DeleteResult, error) {
	opts := options.FindOptions{}
	opts.SetLimit(nRows)
	docs, err := s.repository.Find(collectionName, bson.M{}, &opts)
	if err != nil {
		return nil, err
	}
	var ids []interface{}
	for _, data := range *docs {
		ids = append(ids, data["_id"])
	}

	delopts := options.DeleteOptions{}
	return s.repository.DeleteMany(collectionName, bson.M{"_id": bson.M{"$in": ids } }, &delopts)
}


func (s *service) InsertOne(id string) (*mongo.InsertOneResult, error) {
	opts := options.InsertOneOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.InsertOne("",filter.Filter, &opts)
}

func (s *service) Insert(id string) (*mongo.InsertManyResult, error) {
	opts := options.InsertManyOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	var a []interface{}
	a[0] = filter.Filter
	return s.repository.InsertMany("", a, &opts)
}

func (s *service) UpdateOne(id string) (*mongo.UpdateResult, error) {
	opts := options.UpdateOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.UpdateOne("", filter.Filter, "", &opts)
}

func (s *service) Update(id string) (*mongo.UpdateResult, error) {
	opts := options.UpdateOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.UpdateMany("", filter.Filter, "", &opts)
}

type MongoId struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
}