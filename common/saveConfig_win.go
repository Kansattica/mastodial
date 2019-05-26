// +build windows

package common

import (
	"os"
	"syscall"
)

func osGetConfigFile() (*os.File, error) {
	return os.OpenFile(ConfigLocation, os.O_RDWR|os.O_CREATE|syscall.FILE_SHARE_DELETE, 0644)
}
