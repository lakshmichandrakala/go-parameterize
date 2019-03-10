package go_parameterize

import (
	"regexp"
	"strings"
)

func Parameterize(in string) string {
	reAlphaNum := regexp.MustCompile("[^A-Za-z0-9]+")
	out := reAlphaNum.ReplaceAllString(in, "-")
	out = strings.Trim(out, "-")
	return strings.ToLower(out)
}
