package repository

// A Platform represents information that spans across links such a Spotify title and logo url. It does not have,
// anything to do specifically with the music player, but just the platform info for an embedded resource, i.e It could
// represent Amazon if we wanted an affiliate book link type.
type Platform struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Url     string `db:"url"`
	LogoUrl string `db:"logo_url"`
}
