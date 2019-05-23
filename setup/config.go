package setup

import (
	"fmt"
	"github.com/kansattica/mastodial/common"
	"os"
)

func hasforce(arr []string) (without []string, force bool) {
	flag := -1
	for i, val := range arr {
		if val == "-f" || val == "--force" {
			flag = i
			break
		}
	}

	if flag == -1 {
		return arr, false
	}

	return append(arr[:flag], arr[flag+1:]...), true
}

func help(arr []string) bool {
	for _, val := range arr {
		if val == "-h" {
			return true
		}
	}

	return false
}

func config() {
	args, force := hasforce(os.Args)

	if help(args) || len(args) <= 3 {
		valid()
		return
	}

	toset := args[3]

	var value string
	if len(args) > 4 {
		value = args[4]
	} else {
		val, err := common.GetConfig(toset)
		if err != nil {
			fmt.Printf("for key %s: error %s\n", toset, err)
		}
		fmt.Println(val)
		return
	}

	err := common.SetConfig(toset, value, force)

	if err != nil {
		fmt.Println("could not set key %s: %s", toset, err)
	}

	return
}

func valid() {
	fmt.Printf("%s set config [optionname] [optionvalue] - set values in config file\n", os.Args[0])
	fmt.Printf("%s set, get, and setup are all the same, use whichever works best for you.\n", os.Args[0])
	fmt.Printf("This does mean that %s get can change your config file if you give it both arguments.\n", os.Args[0])
	fmt.Println("This is meant to help you set up your config file- feel free to edit " + common.ConfigLocation + " by hand if you prefer.")
	fmt.Println("Using this tool will create a " + common.ConfigLocation + ".bak file with the value of your " + common.ConfigLocation + " before the change.")
	fmt.Println("Omitting [optionvalue] will print the current value for that option.")
	fmt.Println("Option names are case insensitive. InstanceUrl, Instanceurl, and instanceurl all work.")
	fmt.Println("Keep your config file safe! Anyone with the file can post to your Mastodon account.")
	fmt.Println("Available options:")
	for _, val := range common.Alloptions {
		fmt.Println(val)
	}
}
