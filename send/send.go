package send

import (
	"fmt"

	"github.com/kansattica/mastodial/common"
)

func Send(args []string) {
	if len(args) == 1 {
		usage()
		return
	}

	fmt.Println(args)
	acts, err := parseArgsToActions(args[1:])

	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("%+v\n", acts)
	fmt.Println(acts)
}

type ad struct {
	action, args, description string
}

var actiondescriptions = [...]ad{
	{"nop", "(none)", "do nothing"},
	{"fav", "[pid]*", "favorites the post with post id [pid]. You can specify multiple post IDs, separated with spaces."},
	{"boost", "[pid]*", "boosts the post with post id [pid]. You can specify multiple post IDs, separated with spaces."},
	{"post", "[text]", "creates a new post with the text in [text]."},
	{"post", "[cw] | [text]", "creates a new post with the text in [text] and the content warning [cw]. "},
	{"reply", "[to] | [text]", "reply to post id [to] with the text in [text]."},
	{"reply", "[to] | [cw] | [text]", "reply to post id [to] with the text in [text] and the content warning [cw]."},
}

func usage() {
	fmt.Println("Mastodial - Mastodon (and compatible) client for low-bandwidth connections.", "Usage:")
	fmt.Printf("\t%s send [action] [args] - performs [action] immediately\n", common.CommandName)
	fmt.Printf("\t%s -q send [action] [args] - saves action to a queue file (%s, settable with -qf) to be sent later\n", common.CommandName, common.QueueLocation)
	fmt.Println("Notice that flags have to be immediately after %s! They won't work anywhere else.", common.CommandName)
	fmt.Println("The valid [action]s and their [args] are as follows:")
	for _, val := range actiondescriptions {
		fmt.Println("\t", val.action, "\t", val.args, "\t", val.description)
	}
	fmt.Println("The pipe character | is used to separate arguments. If you want to post a pipe character, put another pipe before it, like this: ||")
}
