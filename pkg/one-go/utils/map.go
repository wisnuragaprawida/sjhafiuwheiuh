package utils

import "strings"

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))

	for _, s := range slice {
		set[strings.ToLower(s)] = struct{}{}
	}
	_, ok := set[strings.ToLower(item)]
	return ok
}
