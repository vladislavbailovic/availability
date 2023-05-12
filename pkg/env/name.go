package env

type Name uint

const (
	SiteID Name = iota
	SiteURL
	PreviouslyDown

	DBConnURI

	TotalNamesCount
)

func (x Name) String() string {
	switch x {
	case SiteID:
		return "AVBL_SITE_ID"
	case SiteURL:
		return "AVBL_SITE_URL"
	case PreviouslyDown:
		return "AVBL_PREVIOUSLY_DOWN"

	case DBConnURI:
		return "AVBL_DBCONN_URI"

	default:
		panic("unknown env var")
	}
}
