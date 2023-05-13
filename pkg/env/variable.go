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

	ApiPortData
	ApiSecretData

	TotalNamesCount
)

func (x Variable) String() string {
	name, ok := knownEnvVars[x]
	if !ok {
		panic("unknown env var")
	}
	return name
}

func (x Variable) Value() string {
	return os.Getenv(x.String())
}

func (x Variable) Expect() string {
	val := x.Value()
	if val == "" {
		panic(fmt.Sprintf("missing required env var: %q", x.String()))
	}
	return val
}

func (x Variable) WithFallback(fb string) string {
	if val := x.Value(); val != "" {
		return val
	}
	return fb
}
