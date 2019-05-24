package setup

import (
	"fmt"
	"github.com/kansattica/mastodial/common"
)

func init() {

}

func Setup(args []string) {
	if len(args) < 2 {
		usage()
		return
	}

	switch args[1] {
	case "app":
		app(args)
	case "creds":
		creds(args)
	case "config":
		config(args)
	default:
		usage()
	}

}

func usage() {
	fmt.Println("Mastodial - Mastodon (and compatible) client for low-bandwidth connections.", "Usage:")
	fmt.Printf("\t%s setup - This help text\n", common.CommandName)
	fmt.Printf("\t%s setup app - Register an application with Mastodon (do this first)\n", common.CommandName)
	fmt.Printf("\t%s setup creds - Set up user credentials for your account (do this second)\n", common.CommandName)
	fmt.Printf("\t%s set config - Manage your config file, including setting your instance URL.\n", common.CommandName)
	fmt.Println("'get', 'set' and 'setup' are interchangable. Run a subcommand without arguments for help and usage information.")
}
