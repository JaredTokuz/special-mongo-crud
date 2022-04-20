package crud

import (
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Service is an interface from which our api module can access our repository of all our models
type Service interface {
	FindOne(collection string, id string) (*interface{}, error)
	Find(collection string, id string) (*[]interface{}, error)
	ExpireDocs(collection string, field string, days string) (*interface{}, error)
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

func (s *service) FindOne(collection string, id string) (*interface{}, error) {
	opts := options.FindOneOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.FindOne(collection, filter.Filter, &opts)
}

func (s *service) Find(collection string, id string) (*[]interface{}, error) {
	opts := options.FindOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.Find(collection, filter.Filter, &opts)
}

func (s *service) ExpireDocs(collection string, field string, days string) (*interface{}, error) {
	opts := options.DeleteOptions{}
	d, _ := strconv.Atoi(days)
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0,0,0,0, now.Location())
	adjdate := date.AddDate(0,0,-d)
	return s.repository.DeleteMany(collection, bson.M{field: bson.M{"$lt": adjdate}}, &opts)
}