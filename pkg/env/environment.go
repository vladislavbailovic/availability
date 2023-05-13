package env

import (
	"fmt"
	"os"
)

func Expect(what Name) string {
	val := os.Getenv(what.String())
	if val == "" {
		panic(fmt.Sprintf("missing required env var: %q", what.String()))
	}
	return val
}
