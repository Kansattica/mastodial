package send

import "fmt"
import "github.com/kansattica/mastodial/common"

func processqueue(todo []action) error {

	body := make(map[string]string)
	for _, act := range todo {
		endpoint := ""
		if body == nil || len(body) > 0 {
			body = make(map[string]string)
		}
		method := "POST"
		switch act.Act {
		case Fav:
			endpoint = fmt.Sprintf("/api/v1/statuses/%s/favourite", act.PostId)
		case Boost:
			endpoint = fmt.Sprintf("/api/v1/statuses/%s/reblog", act.PostId)
		case Del:
			endpoint = "/api/v1/statuses/" + act.PostId
			method = "DELETE"
		case Reply:
			body["in_reply_to_id"] = act.PostId
			fallthrough
		case Post:
			body["status"] = act.Text
			body["visibility"] = "public"
			body["spoiler_text"] = act.CW
			endpoint = "/api/v1/statuses"
		}
		resp, err := common.MakeAuthenticatedRequest(endpoint, method, body, nil)
		if resp.StatusCode == 408 {
			fmt.Println("Your request timed out, but the action may still have gone through.")
		}
		if err != nil {
			fmt.Println("Got an error back. Continuing. Detailed error:", err)
		}
		parsed, err := common.ParseBody(resp.Body)
		if err != nil {
			fmt.Println("Couldn't parse response body. Continuing. Detailed error:", err)
		} else {
			if _, prs := parsed["id"]; !prs {
				parsed["id"] = act.PostId
			}
			fmt.Printf("Successfully %sed post ID %s (content: %s)\n", act.Act, parsed["id"], parsed["content"])
		}

	}

	return nil
}
