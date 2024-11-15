package fp

import "strings"

func join(pattern []string) string {
	return strings.Join(pattern, ",")
}

func split(pattern string) []string {
	return strings.Split(pattern, ",")
}
