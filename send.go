package main

import "flag"
import "fmt"

type sendoptions struct {
	Send bool
	FilePath string
	ToDo Action
	PostId string
}

var opt sendoptions

func init() {
	const (
		sendUsage = "When true, send all queued calls before finishing."
		fileUsage = "Operate on the specified file."
		shorthand = " (shorthand)"
	)
	flag.BoolVar(&opt.Send, "send", false, sendUsage)
	flag.BoolVar(&opt.Send, "s", false, sendUsage + shorthand)

	flag.StringVar(&opt.FilePath, "file", "queue", fileUsage)
	flag.StringVar(&opt.FilePath, "f", "queue", fileUsage + shorthand)
}

func Send() {
	flag.Parse()

	fmt.Println("%+v\n", opt)
	fmt.Println(post.String())
}
