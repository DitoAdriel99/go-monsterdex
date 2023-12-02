package rules

import (
	"regexp"
	"strings"
)

var base64Regex = regexp.MustCompile("^(?:[A-Za-z0-9+\\\\/]{4})*(?:[A-Za-z0-9+\\\\/]{2}==|[A-Za-z0-9+\\\\/]{3}=|[A-Za-z0-9+\\\\/]{4})$")

func IsValidBase64(in string) bool {
	parts := strings.Split(in, ",")
	if len(parts) == 2 {
		in = parts[1]
	}

	return base64Regex.MatchString(in)
}

func IsAllowedImageExtension(extenstion string) bool {

	allowedExtensions := []string{"jpeg", "jpg", "png"}

	for _, allowedExt := range allowedExtensions {
		if extenstion == allowedExt {
			return true
		}
	}

	return false
}
