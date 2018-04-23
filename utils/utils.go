package utils

import "strings"

// Is the string in a simple array
func InArray(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Extract the string after the Bearer
func ExtractTokenFromBearerHeader(raw string) string {
	accessToken := strings.SplitAfter(raw, "Bearer ")[1]
	return accessToken
}
