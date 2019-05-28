package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/grokify/html-strip-tags-go"
)

func MakeAuthenticatedRequest(endpoint, method string, body, queryParams map[string]string) (resp *http.Response, err error) {
	tok := GetConfig(AccessToken)
	return MakeRequest(endpoint, method, body, queryParams, map[string]string{
		"Authorization": "Bearer " + tok,
	})
}

func MakeRequest(endpoint, method string, body, queryParams, header map[string]string) (resp *http.Response, err error) {
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
	fmt.Printf("%s\n", bodyjson)

	if err != nil {
		fmt.Println("Failed to serialize request. This probably isn't your fault. json.Marshal said: " + err.Error())
		return
	}

	fmt.Printf("POST %s (Sending %d bytes)\n", iurl.String(), len(bodyjson))
	//return http.Post(iurl.String(), "application/json", bytes.NewReader(bodyjson))

	client := &http.Client{}
	req, err := http.NewRequest(method, iurl.String(), bytes.NewReader(bodyjson))

	if err != nil {
		return
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json;charset=UTF-8")

	return client.Do(req)
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

	if b, shouldClose := body.(io.Closer); shouldClose {
		b.Close()
	}

	if val, prs := resp["content"]; prs {
		resp["content"] = html.UnescapeString(strip.StripTags(val.(string)))
	} else {
		resp["content"] = "(no body)"
	}

	return

}
