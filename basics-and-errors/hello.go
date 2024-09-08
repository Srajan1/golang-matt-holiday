package hello

import (
	"strings"
)

func Say(names []string) string {
	if len(names) > 0 {
		return "Hello " + strings.Join(names, ", ") + "!"
	} else {
		return "Hello Anonymous!"
	}
}
