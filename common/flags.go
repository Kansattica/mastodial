package common

import "flag"

var CommandName string //gets set in Main

var ConfigLocation string
var Force bool
var QueueRequests bool
var QueueLocation string
var ReadStdin bool

const (
	sh             = " (shorthand)"
	configusage    = "Use the specified config file."
	forceusage     = "Try to write new files, even if creating a backup failed."
	queueusage     = "Queue requests to Mastdon to a file for later. Only works with the recieve subcommand."
	qlocationusage = "File to store or read queued requests."
	stdinusage     = "When giving actions to send, read from stdin instead of the shell."
)

func init() {
	flag.StringVar(&ConfigLocation, "config", "mdconfig.json", configusage)
	flag.StringVar(&ConfigLocation, "c", "mdconfig.json", configusage+sh)

	flag.BoolVar(&Force, "force", false, forceusage)
	flag.BoolVar(&Force, "f", false, forceusage+sh)

	flag.BoolVar(&QueueRequests, "queue", false, queueusage)
	flag.BoolVar(&QueueRequests, "q", false, queueusage+sh)

	flag.StringVar(&QueueLocation, "queuefile", "mdqueue.json", qlocationusage)
	flag.StringVar(&QueueLocation, "qf", "mdqueue.json", qlocationusage+sh)

	flag.BoolVar(&ReadStdin, "stdin", false, stdinusage)

}
