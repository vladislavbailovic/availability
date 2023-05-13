package env

var knownEnvVars map[Variable]string = map[Variable]string{
	SiteID:         "AVBL_SITE_ID",
	SiteURL:        "AVBL_SITE_URL",
	PreviouslyDown: "AVBL_PREVIOUSLY_DOWN",
	DBConnURI:      "AVBL_DBCONN_URI",
	ApiPortCNC:     "AVBL_API_PORT_CNC",
	ApiSecretCNC:   "AVBL_API_SECRET_CNC",
}
