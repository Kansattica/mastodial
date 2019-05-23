package common

import (
	"errors"
	"net/url"
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

		if given[len(given)-1:] != "/" { //if the last character isn't a slash, add a trailing slash
			given += "/"
		}

		return given, nil
	},
}

func isvalid(key, val string) (string, error) {
	strat, prs := strategies[key]

	if !prs {
		return val, nil
	}

	return strat(val)
}
