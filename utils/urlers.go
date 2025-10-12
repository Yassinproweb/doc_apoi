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
