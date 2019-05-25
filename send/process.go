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
		switch act.Act {
		case Fav:
			endpoint = fmt.Sprintf("/api/v1/statuses/%s/favourite", act.PostId)
		case Boost:
			endpoint = fmt.Sprintf("/api/v1/statuses/%s/reblog", act.PostId)
		case Reply:
			body["in_reply_to_id"] = act.PostId
			fallthrough
		case Post:
			body["status"] = act.Text
			body["visibility"] = "public"
			body["spoiler_text"] = act.CW
			endpoint = "/api/v1/statuses"
		}
		resp, err := common.MakeAuthenticatedPost(endpoint, body, nil)
		if err != nil {
			fmt.Printf("Had a problem with: %+v. Continuing. Detailed error: %s", err)
		}
		parsed, err := common.ParseBody(resp.Body)
		if err != nil {
			fmt.Printf("Couldn't parse response body. Continuing. Detailed error: %s", err)
		}

		fmt.Printf("Successfully %sed post ID %s (content: %s)", act.Act, parsed["id"], parsed["content"])
	}

	return nil
}
