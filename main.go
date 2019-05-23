package main

import (
	"fmt"
	"os"

	"github.com/kansattica/mastodial/recv"
	"github.com/kansattica/mastodial/send"
	"github.com/kansattica/mastodial/setup"
)

func main() {
	if (len(os.Args) == 1) {
		usage()
		return
	}

	switch os.Args[1] {
	case "send":
		send.Send()
	case "recv", "receive":
		recv.Recv()
	case "setup", "set", "get":
		setup.Setup()
	default:
		usage()
	}

}

func usage() {
	cmd := os.Args[0]
	fmt.Println("Mastodial - Mastodon (and compatible) client for low-bandwidth connections.", "Usage:")
	fmt.Printf("\t%s recv  - recieve posts\n", cmd)
	fmt.Printf("\t%s send  - send posts\n", cmd)
	fmt.Printf("\t%s setup - set up connection, set options\n", cmd)
	fmt.Println("Append any subcommand with -h for help and usage information.")
}
