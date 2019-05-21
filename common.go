package main

import "errors"

type Action int

const (
	nop Action = iota
	post
	fav
	boost
	reply
)

var actstr = map[string]Action{
	"nop":   nop,
	"post":  post,
	"fav":   fav,
	"boost": boost,
	"reply": reply,
}

func ParseAction(str string) (Action, error) {
	val, prs := actstr[str]

	if prs {
		return val, nil
	} else {
		return nop, errors.New("Invalid opcode " + str)
	}
}
