package main

import "os"
import "fmt"

func main() {

	if len(os.Args) == 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "send":
		Send()
	case "recv":
	case "receive":
		Recv()
	}

}

func usage() {
	cmd := os.Args[0]
	fmt.Println("Mastodial - Mastodon (and compatible) client for low-bandwidth connections.", "Usage:")
	fmt.Printf("\t%s recv  - recieve posts\n", cmd)
	fmt.Printf("\t%s send  - send posts\n", cmd)
	fmt.Printf("\t%s setup - set up connection, set options\n", cmd)
	fmt.Println("Append any command with -h for help and usage information.")
}
