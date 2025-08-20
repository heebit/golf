package utils

import (
	"regexp"
	"strings"
)

func SanitizeFileName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "_")

	re := regexp.MustCompile(`[^a-zA-Z0-9а-яА-Я_-]`)
	name = re.ReplaceAllString(name, "")

	return name
}
