package db

import (
	"context"
	"errors"

	"github.com/jdebes/LinkTreeBackend/db/repository"
)

type mockStore string

var (
	mockKey = mockStore("mockStore")
)

type MockStore struct {
	linkStore     map[int64]*repository.Link
	platformStore map[int64]*repository.Platform
	userStore     map[int64]*repository.User
	nextId        int64
}

func (s *MockStore) NextId() int64 {
	next := s.nextId
	s.nextId++
	return next
}

func NewMockStore() *MockStore {
	store := MockStore{
		linkStore:     make(map[int64]*repository.Link),
		platformStore: make(map[int64]*repository.Platform),
		userStore:     make(map[int64]*repository.User),
		nextId:        0,
	}

	mockUser := repository.User{
		ID: store.NextId(),
	}
	store.userStore[mockUser.ID] = &mockUser

	mockPlatform := repository.Platform{
		ID:      store.NextId(),
		Name:    "Spotify",
		Url:     "https://www.spotify.com/",
		LogoUrl: "https://www.spotify.com/logo.png",
	}
	store.platformStore[mockPlatform.ID] = &mockPlatform

	return &store
}

func (s MockStore) Links() []*repository.Link {
	links := make([]*repository.Link, 0, len(s.linkStore))
	for _, value := range s.linkStore {
		links = append(links, value)
	}

	return links
}

func WithStore(ctx context.Context, db *MockStore) context.Context {
	return context.WithValue(ctx, key, db)
}

func Store(ctx context.Context) (*MockStore, error) {
	if db, ok := ctx.Value(key).(*MockStore); ok {
		return db, nil
	}

	return nil, errors.New("No DB")
}
