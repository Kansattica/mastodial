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
	todo, err := ParseActionType(args[0])

	if err != nil {
		return
	}

	args = args[1:]

	var thisact action
	thisact.Act = todo
	minargs := 1
	fields := []*string{&thisact.Text, &thisact.CW}
	switch todo {
	case Fav, Boost, Del:
		for _, str := range args {
			for _, val := range strings.Fields(str) {
				acts = append(acts, action{Act: todo, PostId: val})
			}
		}

	case Reply:
		fields = []*string{&thisact.PostId, &thisact.Text, &thisact.CW}
		minargs = 2
		fallthrough
	case Post:
		parsePercentArgs(args, fields)
		if thisact.Text == "" {
			err = fmt.Errorf("You have to give some text for %s", todo)
			return
		}
		acts = append(acts, thisact)
	}

	if len(args) == 0 {
		err = fmt.Errorf("You need %d or more arguments.", minargs)
		return
	}

	return
}

func parsePercentArgs(args []string, fields []*string) {
	build := ""
	fieldcount := 0
	for _, val := range args {
		if fieldcount >= len(fields) {
			return
		}

		switch val {
		case "%":
			*(fields[fieldcount]) = strings.Replace(build, "%%", "%", -1)
			build = ""
			fieldcount++
		default:
			build = build + val + " "
		}
	}
	*(fields[fieldcount]) = strings.TrimRight(strings.Replace(build, "%%", "%", -1), "\t \r\n")
}

type ActionType int

const (
	Nop ActionType = iota
	Post
	Fav
	Boost
	Reply
	Del
)

var actstr = map[string]ActionType{
	"nop":    Nop,
	"post":   Post,
	"fav":    Fav,
	"boost":  Boost,
	"reply":  Reply,
	"delete": Del,
}

func ParseActionType(str string) (ActionType, error) {
	val, prs := actstr[strings.ToLower(str)]

	if prs {
		return val, nil
	} else {
		return Nop, fmt.Errorf("Invalid action %s", str)
	}
}
