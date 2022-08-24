package utils

import "strings"

func Slugify(name string) string {
	name = strings.ToLower(name)
	arr := strings.Split(name, " ")
	slug := strings.Join(arr, "-")
	return slug
}
