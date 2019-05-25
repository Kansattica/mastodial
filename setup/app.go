package setup

import (
	"fmt"
	"github.com/kansattica/mastodial/common"
)

func app(args []string) {

	url := common.GetConfig(common.InstanceUrl)

	if url == "" {
		fmt.Printf("Please set your instance URL by running:\n %s setup set config instanceurl https://[your instance url]\n", common.CommandName)
		return
	}

	if common.OptionExists(common.ClientId) && common.OptionExists(common.ClientSecret) {
		fmt.Println("You have already registered an app with", url, ". If this is for a new instance, make a new config file with the -c flag or delete", common.ConfigLocation)
		fmt.Printf("You probably want to do this next:\n %s setup creds\n", common.CommandName)
		return
	}

	resp, err := common.MakePostRequest("/api/v1/apps", map[string]string{
		"client_name":   "Mastodial",
		"redirect_uris": "urn:ietf:wg:oauth:2.0:oob",
		"scopes":        "read write",
		"website":       "https://github.com/Kansattica/mastodial",
	}, nil, nil)

	if err != nil {
		fmt.Println("Failed to register app. Please try again. Error: " + err.Error())
	}

	body, err := common.ParseBody(resp.Body)

	if err != nil {
		fmt.Println("Failed to register app. Please try again. Error: " + err.Error())
		return
	}

	fmt.Printf("%+v\n", body)
	common.SetConfig(common.ClientId, body["client_id"].(string), true)
	err = common.SetConfig(common.ClientSecret, body["client_secret"].(string), false)

	if err != nil {
		fmt.Println("Failed to set config. Please try again. Error: ", err.Error())
	}

	fmt.Println("Registered app!")

}
