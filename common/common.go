package common

import "errors"
import "strings"

type Action int

const (
	Nop Action = iota
	Post
	Fav
	Boost
	Reply
)

var actstr = map[string]Action{
	"nop":   Nop,
	"post":  Post,
	"fav":   Fav,
	"boost": Boost,
	"reply": Reply,
}

func ParseAction(str string) (Action, error) {
	val, prs := actstr[strings.ToLower(str)]

	if prs {
		return val, nil
	} else {
		return Nop, errors.New("Invalid opcode " + str)
	}
}
