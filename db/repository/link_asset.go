package repository

type AssetType int

func (a AssetType) String() string {
	return [...]string{"Resource", "Embedded", "Integration"}[a]
}

const (
	// Resource is an asset that we redirect to
	Resource AssetType = iota
	// Embedded is something we will display in an iframe
	Embedded
	// Integration is information from an external source that we cant redirect to or embed directly into the page,
	// thus we will need to call it and get what we are after then return it.
	Integration
)

// A LinkAsset represents a resource that a link may show. A Link may have multiple resources depending on its type.
// Since we could have many potential types, I have abstracted this to a single Link (i.e music type) containing multiple
// LinkAssets (an embeddable link which represents a song, playlist, album) etc.
// The thinking behind this is that any link type is just a collection of external resources to displayed.
type LinkAsset struct {
	ID       int64     `db:"id"`
	Type     AssetType `db:"type"`
	Name     string    `db:"name"`
	Url      string    `db:"url"`
	Platform `db:"platform"`
}
