package setup

import (
	"errors"
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

	authInfo := make(map[string]string)
	var granttype string

	switch strat {
	case authcode:
		granttype = "authorization_code"
		authInfo["code"] = common.GetConfig(common.AuthCode)
	case userpass:
		granttype = "client_credentials"
		authInfo["username"] = common.GetConfig(common.Username)
		authInfo["password"] = common.GetConfig(common.Password)
	case nocreds:
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
		return
	}

	atoken, err := getToken(granttype, authInfo)

	if err != nil {
		fmt.Println(err)
		if strat == authcode {
			fmt.Println("If you need a new auth code, you can get one here:")
			fmt.Println(authcodeurl())
			fmt.Println("If you'd prefer to try username and password auth instead, remove your expired auth code with this command:")
			fmt.Printf("%s setup config authcode \"\"\n", common.CommandName)
		}
		return
	}

	aerr := common.SetConfig(common.AccessToken, atoken, false)

	if aerr != nil {
		fmt.Println("Failed to save tokens. SetConfig said:\n", aerr)
	}

	fmt.Println("Got your access token! You should be ready to go.")

	if strat == authcode {
		fmt.Println("Deleting your used authorization code.")
		common.DeleteConfig(common.AuthCode)
	}
	return

}

func getToken(granttype string, authInfo map[string]string) (string, error) {
	authInfo["grant_type"] = granttype
	authInfo["client_id"] = common.GetConfig(common.ClientId)
	authInfo["client_secret"] = common.GetConfig(common.ClientSecret)
	authInfo["redirect_uri"] = "urn:ietf:wg:oauth:2.0:oob"
	resp, err := common.MakePostRequest("/oauth/token", authInfo, nil, nil)

	if err != nil {
		return "", errors.New("Could not get token. Tried authentication grant type: " + granttype + " Post request returned: " + err.Error())
	}

	body, err := common.ParseBody(resp.Body)

	if err != nil {
		return "", errors.New("Could not parse response. Tried authentication grant type: " + granttype + " Parser returned: " + err.Error())
	}

	atoken, aok := body["access_token"].(string)
	if aok {
		return atoken, nil
	}

	return "", fmt.Errorf("Did not get tokens back. Instead got: %v", body)

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
