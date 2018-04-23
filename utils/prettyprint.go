package utils

import (
	"encoding/json"
	"log"
)

// Dump to log the interface as a pretty json
func PrettyPrintJSON(i interface{}, indent string) {
	b, err := json.MarshalIndent(i, "", indent)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(b))
}
