package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kansattica/mastodial/common"
	"github.com/kansattica/mastodial/recv"
	"github.com/kansattica/mastodial/send"
	"github.com/kansattica/mastodial/setup"
)

func main() {
	flag.Parse()
	args := flag.Args()
	common.CommandName = os.Args[0]
	if len(args) == 0 {
		usage(common.CommandName)
		return
	}

	fmt.Println(args)
	switch args[0] {
	case "send":
		send.Send(args)
	case "recv", "receive":
		recv.Recv(args)
	case "setup", "set", "get":
		setup.Setup(args)
	default:
		usage(common.CommandName)
	}

}

func usage(cmd string) {
	fmt.Println("Mastodial - Mastodon (and compatible) client for low-bandwidth connections.", "Usage:")
	fmt.Printf("\t%s recv  - recieve posts\n", cmd)
	fmt.Printf("\t%s send  - send posts\n", cmd)
	fmt.Printf("\t%s setup - set up connection, set options\n", cmd)
	fmt.Println("Append any subcommand with -h for help and usage information.")
}
