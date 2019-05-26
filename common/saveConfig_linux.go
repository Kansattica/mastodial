// +build linux

package common

import (
	"os"
)

func osGetConfigFile() (*os.File, error) {
	return os.OpenFile(ConfigLocation, os.O_RDWR|os.O_CREATE, 0644)
}
