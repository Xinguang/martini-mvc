package helpers

import (
	"regexp"
)

type ExRegexp string

func (s ExRegexp) replace(input string, repl string) string {
	re := regexp.MustCompile(string(s))
	return re.ReplaceAllString(input, repl)
}

func (s ExRegexp) match(input string) bool {
	matched, err := regexp.MatchString(string(s), input)
	if err == nil {
		return matched
	} else {
		return false
	}
}

func (s ExRegexp) find(input string) string {
	re := regexp.MustCompile(string(s))
	return re.FindString(input)
}

/*
(?P<first>[A-Z][a-z]+)(?P<last>[A-Z][a-z]+)
([A-Z][a-z]+)([A-Z][a-z]+)
*/
func (s ExRegexp) findAll(input string) map[string]string {
	re := regexp.MustCompile(string(s))

	matches := re.FindStringSubmatch(input)
	if len(matches) > 0 && matches[0] == input {
		params := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if len(name) > 0 {
				params[name] = matches[i]
			} else {
				params[string(i)] = matches[i]
			}
		}
		return params
	}
	return nil
}
