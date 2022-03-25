package post

import (
	"context"
	"fmt"
)

// Post - a representation of the comment structure for our service
type Post struct {
	ID string
	Slug string
	Author string
	Body string
}

type Store interface {
	GetPost(context.Context, string) (Post, error)
}

// Service - is the struct on which all our logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetPost(ctx context.Context, id string) (Post, error) {
	fmt.Println("retrieving post")
	return Post{}, nil
}