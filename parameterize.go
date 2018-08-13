package go_parameterize

import (
	"regexp"
	"strings"
)

func Parameterize(in string) string {
	reAlphaNum := regexp.MustCompile("[^A-Za-z0-9]+")
	reTrim := regexp.MustCompile("^-|-$")

	out := reAlphaNum.ReplaceAllString(in, "-")
	out = reTrim.ReplaceAllString(out, "")

	return strings.ToLower(out)
}
