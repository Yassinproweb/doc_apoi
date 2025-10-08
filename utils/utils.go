package utils

import "slices"

func IsAllowedFileType(contentType string, allowedTypes []string) bool {
	return slices.Contains(allowedTypes, contentType)
}
