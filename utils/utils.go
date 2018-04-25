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
	accessToken := strings.SplitAfter(raw, "Bearer ")[1]
	return accessToken
}

func DumpResponseBody(resp *http.Response) {
	fmt.Println(resp.StatusCode)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
