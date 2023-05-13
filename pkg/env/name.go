package env

type Name uint

const (
	SiteID Name = iota
	SiteURL
	PreviouslyDown

	DBConnURI

	ApiPortCNC
	ApiSecretCNC

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

	case ApiPortCNC:
		return "AVBL_API_PORT_CNC"
	case ApiSecretCNC:
		return "AVBL_API_SECRET_CNC"

	default:
		panic("unknown env var")
	}
}
