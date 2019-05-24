package setup

import (
	"fmt"
	"net/url"

	"github.com/kansattica/mastodial/common"
)

type strategy int

const (
	nocreds  strategy = 0
	authcode          = 1
	userpass          = 2
)

func creds(args []string) {

	strat := pickauth()

	if strat == nocreds {
		fmt.Println("You have two options to set up your credentials. You can either store your username and password, or you can generate an authentication code.")
		fmt.Println("If you can open up a browser, I suggest the authentication code method. Visit the following url and complete the prompts:")
		fmt.Println(authcodeurl())
		fmt.Println("Copy the authentication code from there and run:")
		fmt.Println(common.CommandName, "set config", common.AuthCode, "(value from web page)")
		fmt.Println("Then run this command again.")
		fmt.Println("If you'd rather just use your username and password, run the following two commands:")
		fmt.Println(common.CommandName, "set config", common.Username, "(your username)")
		fmt.Println(common.CommandName, "set config", common.Password, "(your password)")
		fmt.Println("Then run this command again.")
	}

}

func pickauth() strategy {
	val := common.GetConfig(common.AuthCode)

	if val != "" {
		return authcode
	}

	uname := common.GetConfig(common.Username)
	pass := common.GetConfig(common.Password)

	if uname != "" && pass != "" {
		return userpass
	}

	return nocreds
}

func authcodeurl() string {
	rawurl := common.GetConfig(common.InstanceUrl)

	purl, _ := url.Parse(rawurl)
	purl.Path = "/oauth/authorize"

	query := purl.Query()

	query.Add("response_type", "code")
	query.Add("client_id", common.GetConfig(common.ClientId))
	query.Add("client_secret", common.GetConfig(common.ClientSecret))
	query.Add("redirect_uri", "urn:ietf:wg:oauth:2.0:oob")
	query.Add("scope", "read write")

	purl.RawQuery = query.Encode()
	return purl.String()
}
