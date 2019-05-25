package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func MakePostRequest(endpoint string, body, queryParams map[string]string) (resp *http.Response, err error) {
	strUrl := GetConfig(InstanceUrl)

	if strUrl == "" {
		fmt.Printf("Please set your instance URL by running:\n %s setup set config instanceurl https://[your instance url]\n", CommandName)
		return nil, err
	}

	iurl, err := url.Parse(strUrl)

	if err != nil {
		fmt.Printf("Could not parse your instance URL: " + strUrl)
		return nil, err
	}

	iurl.Path = endpoint

	q := iurl.Query()

	if queryParams != nil {
		for k, v := range queryParams {
			q.Set(k, v)
		}

		iurl.RawQuery = q.Encode()
	}

	bodyjson, err := json.Marshal(body)

	if err != nil {
		fmt.Println("Failed to serialize request. This probably isn't your fault. json.Marshal said: " + err.Error())
		return
	}

	fmt.Printf("POST %s (Sending %d bytes)\n", iurl.String(), len(bodyjson))
	return http.Post(iurl.String(), "application/json", bytes.NewReader(bodyjson))
}

func ParseBody(body io.Reader) (resp map[string]interface{}, err error) {
	buf := new(bytes.Buffer)

	read, err := buf.ReadFrom(body)

	if err != nil {
		fmt.Printf("Failed to read response after %d bytes. buf.ReadFrom said: %s", read, err.Error())
		return
	}

	err = json.Unmarshal(buf.Bytes(), &resp)

	if err != nil {
		fmt.Printf("Failed to deserialize request. This probably isn't your fault. json.Unmarshal said: %s. Bytes: %d\n", err.Error(), read)
		buf.WriteTo(os.Stdout)
		return
	}

	return

}
