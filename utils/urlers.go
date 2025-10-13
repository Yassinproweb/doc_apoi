package utils

import (
	"fmt"
	"strings"
)

func NormalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	return name
}

func URLer(routex, name string) string {
	res := fmt.Sprintf("/%s/%s", routex, NormalizeName(name))
	return res
}

func Capitalize(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	ws := strings.Fields(s)
	for i, w := range ws {
		ws[i] = strings.ToUpper(string(w[0])) + w[1:]
	}

	return strings.Join(ws, " ")
}
