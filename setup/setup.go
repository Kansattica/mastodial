package setup

import (
	"fmt"
	"os"
)

func init() {

}

func Setup() {
	if len(os.Args) < 3 {
		usage()
		return
	}

	switch os.Args[2] {
	case "app":

	case "creds":

	case "config":
		config()
	default:
		usage()
	}

}

func usage() {
	cmd := os.Args[0]
	fmt.Println("Mastodial - Mastodon (and compatible) client for low-bandwidth connections.", "Usage:")
	fmt.Printf("\t%s setup - This help text\n", cmd)
	fmt.Printf("\t%s setup app - Register an application with Mastodon (do this first)\n", cmd)
	fmt.Printf("\t%s setup creds - Set up user credentials for your account (do this second)\n", cmd)
	fmt.Printf("\t%s set config - Manage your config file, including setting your instance URL.\n", cmd)
	fmt.Println("'get', 'set' and 'setup' are interchangable. Run a subcommand with -h for help and usage information.")
}
