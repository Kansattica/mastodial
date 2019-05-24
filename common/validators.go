package common

import (
	"errors"
	"net/url"
	"strings"
)

type validate func(value string) (string, error)

var strategies = map[string]validate{
	InstanceUrl: func(given string) (string, error) {
		parsed, err := url.Parse(given)

		if err != nil {
			return given, errors.New("Failed to parse URL: " + err.Error())
		}

		if parsed.Scheme != "https" {
			return given, errors.New("URL must start with https://")
		}

		return strings.TrimRight(parsed.String(), "/"), nil
	},
}

func isvalid(key, val string) (string, error) {
	strat, prs := strategies[key]

	if !prs {
		return val, nil
	}

	return strat(val)
}
