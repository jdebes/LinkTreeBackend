package repository

import (
	"context"
	"errors"
)

type mockStore string

var (
	mockKey = mockStore("mockStore")
)

type MockStore struct {
	linkStore     map[int64]*Link
	platformStore map[int64]*Platform
	userStore     map[int64]*User
	nextId        int64
}

func (s *MockStore) NextId() int64 {
	next := s.nextId
	s.nextId++
	return next
}

func NewMockStore() *MockStore {
	store := MockStore{
		linkStore:     make(map[int64]*Link),
		platformStore: make(map[int64]*Platform),
		userStore:     make(map[int64]*User),
		nextId:        0,
	}

	// ID = 0
	mockUser := User{
		ID:    store.NextId(),
		Links: make([]*Link, 0),
	}
	store.userStore[mockUser.ID] = &mockUser

	// ID = 1
	mockPlatform := Platform{
		ID:      store.NextId(),
		Name:    "Spotify",
		Url:     "https://www.spotify.com/",
		LogoUrl: "https://www.spotify.com/logo.png",
	}
	store.platformStore[mockPlatform.ID] = &mockPlatform

	return &store
}

func (s *MockStore) Platform(id int64) (*Platform, error) {
	platform, ok := s.platformStore[id]
	if !ok {
		return nil, errors.New("Platform does not exist")
	}

	return platform, nil
}

func (s *MockStore) User(id int64) (*User, error) {
	user, ok := s.userStore[id]
	if !ok {
		return nil, errors.New("User does not exist")
	}

	return user, nil
}

func (s *MockStore) InsertLink(link *Link, user *User) (*Link, error) {
	// We expect user to have links as the relationship, so this is the quick implementation for that.
	user.Links = append(user.Links, link)
	s.linkStore[link.ID] = link

	return link, nil
}

func WithStore(ctx context.Context, db *MockStore) context.Context {
	return context.WithValue(ctx, mockKey, db)
}

func Store(ctx context.Context) (*MockStore, error) {
	if db, ok := ctx.Value(mockKey).(*MockStore); ok {
		return db, nil
	}

	return nil, errors.New("No DB")
}
