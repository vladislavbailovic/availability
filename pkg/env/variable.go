package env

import (
	"fmt"
	"os"
)

type Variable uint

const (
	SiteID Variable = iota
	SiteURL
	PreviouslyDown

	DBConnURI

	ApiPortCNC
	ApiSecretCNC

	TotalNamesCount
)

func (x Variable) String() string {
	name, ok := knownEnvVars[x]
	if !ok {
		panic("unknown env var")
	}
	return name
}

func (x Variable) Expect() string {
	val := os.Getenv(x.String())
	if val == "" {
		panic(fmt.Sprintf("missing required env var: %q", x.String()))
	}
	return val
}
