package repository

import "time"

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
	ID          int64     `db:"id"`
	Type        LinkType  `db:"type"`
	CreatedDate time.Time `db:"created_date"`
	LinkAssets  []*LinkAsset
}
