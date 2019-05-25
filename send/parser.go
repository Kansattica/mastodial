package send

import (
	"fmt"
	"strings"
)

type action struct {
	Act      ActionType
	PostId   string
	Text, CW string
}

func parseArgsToActions(args []string) (acts []action, err error) {
	fmt.Print(args)
	todo, err := ParseActionType(args[0])

	if err != nil {
		return
	}

	switch todo {
	case Fav, Boost:
		for _, val := range args[1:] {
			acts = append(acts, action{Act: todo, PostId: val})
		}
	}

	return
}

func parseFavBoostArgs(args []string) []string {
	return args[2:]
}

type ActionType int

const (
	Nop ActionType = iota
	Post
	Fav
	Boost
	Reply
)

var actstr = map[string]ActionType{
	"nop":   Nop,
	"post":  Post,
	"fav":   Fav,
	"boost": Boost,
	"reply": Reply,
}

func ParseActionType(str string) (ActionType, error) {
	val, prs := actstr[strings.ToLower(str)]

	if prs {
		return val, nil
	} else {
		return Nop, fmt.Errorf("Invalid action %s", str)
	}
}
