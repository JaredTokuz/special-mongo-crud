package crud

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Service is an interface from which our api module can access our repository of all our models
type Service interface {
	FindOne(id string) (*interface{}, error)
	Find(id string) (*[]interface{}, error)
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

func (s *service) FindOne(id string) (*interface{}, error) {
	opts := options.FindOneOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.FindOne(filter.Filter, &opts)
}

func (s *service) Find(id string) (*[]interface{}, error) {
	opts := options.FindOptions{}
	filter, err := s.repository.FindParameters(id)
	if err != nil {
		return nil, err
	}
	return s.repository.Find(filter.Filter, &opts)
}
