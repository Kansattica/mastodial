package setup

import (
	"fmt"
	"github.com/kansattica/mastodial/common"
)

func app(args []string) {
	resp, err := common.MakePostRequest("/api/v1/apps", map[string]string{
		"client_name":   "Mastodial",
		"redirect_uris": "urn:ietf:wg:oauth:2.0:oob",
		"scopes":        "read write",
		"website":       "https://github.com/Kansattica/mastodial",
	}, nil)

	if err != nil {
		fmt.Println("Failed to register app. Please try again. Error: " + err.Error())
	}

	body, err := common.GetResponse(resp.Body)

	if err != nil {
		fmt.Println("Failed to register app. Please try again. Error: " + err.Error())
		return
	}

	fmt.Printf("%+v\n", body)
	common.SetConfig(common.ClientId, body["client_id"], false, true)
	err = common.SetConfig(common.ClientSecret, body["client_secret"], false, false)

	if err != nil {
		fmt.Println("Failed to set config. Please try again. Error: ", err.Error())
	}

}
