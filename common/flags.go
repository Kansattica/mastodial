package common

import "flag"

var CommandName string //gets set in Main

var ConfigLocation string
var Force bool

const (
	sh          = " (shorthand)"
	configusage = "Use the specified config file."
	forceusage  = "Try to write new files, even if creating a backup failed."
)

func init() {
	flag.StringVar(&ConfigLocation, "config", "mdconfig.json", configusage)
	flag.StringVar(&ConfigLocation, "c", "mdconfig.json", configusage+sh)

	flag.BoolVar(&Force, "force", false, forceusage)
	flag.BoolVar(&Force, "f", false, forceusage+sh)

}
