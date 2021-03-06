package send

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/kansattica/mastodial/common"
)

func Send(args []string) {
	if len(args) == 1 {
		usage()
		return
	}

	args = args[1:]

	if common.ReadStdin {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(ScanSpaceyWords)

		args = args[:1]
		for scanner.Scan() {
			args = append(args, scanner.Text())
		}

		args[len(args)-1] = strings.TrimRight(args[len(args)-1], "\r\n") //Trim trailing newlines, such as from typing your post in and entering ^D

		if scanner.Err() != nil {
			fmt.Println("Error reading from stdin:", scanner.Err())
			return
		}
	}

	acts, err := parseArgsToActions(args)
	sendQueue := acts[0].Act == Queue

	if sendQueue && common.QueueRequests {
		fmt.Println("Hey, you gotta pick one. Do you want to queue stuff or send stuff?")
		fmt.Println("Use the -q option to say that an action should be queued for later.")
		fmt.Println("Use the queue subcommand to execute everything in the queue.")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", acts)

	if sendQueue {
		acts, err = readQueue()

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if common.QueueRequests {
		qacts, err := readQueue()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = saveQueue(append(qacts, acts...))
		fmt.Println("Post queued to", common.QueueLocation)
		return
	} else {
		err = processqueue(acts)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", acts)
	if sendQueue {
		fmt.Println("Queue sent successfully. Renaming to", common.QueueLocation+".sent")
		err = os.Rename(common.QueueLocation, common.QueueLocation+".sent")

		if err != nil {
			fmt.Println("Rename failed. os.Rename said: ", err.Error())
		}

	}

}

type ad struct {
	action, args, description string
}

var actiondescriptions = [...]ad{
	{"nop", "(none)", "do nothing"},
	{"fav", "[pid]*", "favorite the post with post id [pid]. You can list multiple post IDs, separated with spaces."},
	{"boost", "[pid]*", "boost the post with post id [pid]. You can list multiple post IDs, separated with spaces."},
	{"delete", "[pid]*", "delete the post with post id [pid]. You can list multiple post IDs, separated with spaces."},
	{"post", "[text]", "create a new post with the text in [text]."},
	{"post", "[text] % [cw]", "create a new post with the text in [text] and the content warning [cw]. "},
	{"reply", "[to] % [text]", "reply to post id [to] with the text in [text]."},
	{"reply", "[to] % [text] % [cw]", "reply to post id [to] with the text in [text] and the content warning [cw]."},
}

func usage() {
	fmt.Println("Mastodial - Mastodon (and compatible) client for low-bandwidth connections.", "Usage:")
	fmt.Printf("\t%s send - this help text\n", common.CommandName)
	fmt.Printf("\t%s send queue - send everything in the queue\n", common.CommandName)
	fmt.Printf("\t%s send [action] [args] - performs [action] immediately\n", common.CommandName)
	fmt.Printf("\t%s -q send [action] [args] - saves action to a queue file to be sent later\n", common.CommandName)
	fmt.Printf("Notice that flags have to be immediately after %s! They won't work anywhere else.\n", common.CommandName)
	fmt.Println("The queue file is mdqueue.json by default. Specify one with the -df flag.")
	fmt.Println("The valid [action]s and their [args] are as follows:\n")
	for _, val := range actiondescriptions {
		fmt.Println(val.action, "\t", val.args, "\t", val.description)
	}
	fmt.Println("\nThe percent sign character % is used to separate arguments.\nIf you want to post a percent sign, put another percent sign before it, like this: %%")
	fmt.Println("If the shell keeps interpreting your exclamation points and whatnot, try using the -stdin flag. Then, type your post and hit ctrl-D (ctrl-Z on Windows) on its own line or pipe it in from a file.")
}

//copied and adapted from https://golang.org/src/bufio/scan.go?s=13093:13171#L380
func ScanSpaceyWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if r != ' ' {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ' ' {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
