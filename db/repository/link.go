package repository

import (
	"sort"
	"time"

	"github.com/jdebes/LinkTreeBackend/handler/model"
)

type LinkType int

func (a LinkType) String() string {
	return [...]string{"Classic", "MusicPlayer", "ShowsList"}[a]
}

const (
	Classic AssetType = iota
	MusicPlayer
	ShowsList
)

// Link represents a single unit that you would display on a landing page, in the case of a classic link,
// this would be a single title and url. But through the LinkAsset's this could be multiple links which make up
// something more complex like the music player.
type Link struct {
	ID          int64        `db:"id" json:"id"`
	Type        LinkType     `db:"type" json:"type"`
	CreatedDate time.Time    `db:"created_date" json:"createdDate"`
	LinkAssets  []*LinkAsset `json:"linkAssets"`
}

func QueryLinks(store *MockStore, userId int64, orderByCreated bool) ([]*Link, error) {
	user, err := store.User(userId)
	if err != nil {
		return nil, err
	}

	if orderByCreated {
		sort.Slice(user.Links, func(i, j int) bool {
			return user.Links[i].CreatedDate.After(user.Links[j].CreatedDate)
		})
	}

	return user.Links, err
}

func InsertLink(store *MockStore, linkRequest model.Link, userId int64) (*Link, error) {
	user, err := store.User(userId)
	if err != nil {
		return nil, err
	}

	link := Link{
		ID:          store.NextId(),
		Type:        LinkType(linkRequest.Type),
		CreatedDate: time.Now(),
		LinkAssets:  make([]*LinkAsset, 0, len(linkRequest.LinkAssets)),
	}

	for _, value := range linkRequest.LinkAssets {
		var platform *Platform
		if value.PlatformID != nil {
			platform, err = store.Platform(*value.PlatformID)
			if err != nil {
				return nil, err
			}
		}

		asset := LinkAsset{
			ID:       store.NextId(),
			Type:     AssetType(value.Type),
			Name:     value.Name,
			Url:      value.Url,
			Platform: platform,
		}

		link.LinkAssets = append(link.LinkAssets, &asset)
	}

	return store.InsertLink(&link, user)
}
