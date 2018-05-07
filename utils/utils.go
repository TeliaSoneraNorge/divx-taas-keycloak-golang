package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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
	parts := strings.SplitAfter(raw, "Bearer ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func DumpResponseBody(resp *http.Response) {
	fmt.Println(resp.StatusCode)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func HasRefreshTokenExpired(message string) bool {
	return strings.Contains(message, `{"error":"invalid_grant","error_description":"Refresh token expired"}`)
}
