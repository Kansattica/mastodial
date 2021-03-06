package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	InstanceUrl  = "instanceurl"
	ClientId     = "clientid"
	ClientSecret = "clientsecret"
	AccessToken  = "accesstoken"
	Username     = "username"
	Password     = "password"
	AuthCode     = "authcode"
)

var Alloptions = [...]string{InstanceUrl, ClientId, ClientSecret, AccessToken, Username, Password, AuthCode}

var options map[string]string
var ConfigRead bool = false
var triedRead bool = false

func GetConfig(key string) string {
	key = strings.ToLower(key)
	err := iskeygood(key)

	if err != nil {
		return ""
	}

	return options[key]
}

func SetConfig(key string, val string, skipsave bool) error {
	key = strings.ToLower(key)
	err := iskeygood(key)

	if err != nil {
		return err
	}

	val, err = isvalid(key, val)

	if err != nil {
		return errors.New("Validation failed: " + err.Error())
	}

	options[key] = val

	if !skipsave {
		return saveConfig()
	}

	return nil
}

func DeleteConfig(key string) error {
	key = strings.ToLower(key)
	err := iskeygood(key)

	if err != nil {
		return err
	}

	delete(options, key)

	return saveConfig()
}

func OptionExists(key string) bool {
	for _, val := range Alloptions {
		if val == key {
			return true
		}
	}
	return false
}

func OptionSet(key string) bool {
	val, prs := options[key]
	return OptionExists(key) && prs && val != ""
}

func saveConfig() error {
	bytes, err := json.MarshalIndent(options, "", "\t")

	if err != nil {
		fmt.Println("Failed to save changes to disk. Your config files have not been changed. Marshal reported " + err.Error())
		return err
	}

	err = os.Rename(ConfigLocation, ConfigLocation+".bak")

	if !Force && err != nil {
		fmt.Println("Failed to backup your old config file. Your config files have not been changed. Try again with the --force parameter to try saving changes anyway. Rename reported " + err.Error())
		return err
	}

	err = ioutil.WriteFile(ConfigLocation, bytes, 0600)

	if err != nil {
		fmt.Println("Failed to write new config file. Your old file has been saved to " + ConfigLocation + ".bak. Overwrite your current config file as needed. WriteFile reported " + err.Error())
		return err
	}

	return nil

}

func readConfig() {
	if triedRead {
		return
	}

	file, err := os.OpenFile(ConfigLocation, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		fmt.Println("Could not create or open config file. OpenFile reported " + err.Error())
		return
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Could not open config file. ReadAll reported " + err.Error())
		ConfigRead = false
		return
	}

	if len(bytes) > 0 {
		err = json.Unmarshal(bytes, &options)

		if err != nil {
			fmt.Println("Could not parse config file. Unmarshal reported " + err.Error())
			ConfigRead = false
			return
		}
	} else {
		options = make(map[string]string)
	}

	ConfigRead = true
	return
}

func iskeygood(key string) error {
	readConfig()
	if !ConfigRead {
		return errors.New("Could not read config file.")
	}

	if !OptionExists(key) {
		return errors.New("Option does not exist.")
	}

	return nil
}
