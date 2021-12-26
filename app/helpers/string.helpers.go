package helpers

import (
	"log"
	"regexp"
)

func ReplaceSpace(s string) string {
	regex, err := regexp.Compile(" ")
	if err != nil {
		log.Println(err)
		return ""
	}
	s_formatted := regex.ReplaceAllString(s, "-")
	return s_formatted
}
