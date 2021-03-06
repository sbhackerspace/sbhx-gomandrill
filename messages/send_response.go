// Steve Phillips / elimisteve
// 2014.03.29

package messages

import (
	"encoding/json"
	"fmt"
)

type SendResponse struct {
	Email          string `json:"email"`
	Status         string `json:"status"`
	Id             string `json:"_id"`
	RejectedReason string `json:"rejected_reason"`
}

func parseSendResponses(body []byte) ([]*SendResponse, error) {
	var resp []*SendResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("Error parsing *SendResponse: %v\n", err)
	}
	return resp, nil
}

func ValidateSendResponses(body []byte) error {
	resp, err := parseSendResponses(body)
	if err != nil {
		return err
	}

	// Ensure valid status
	for _, r := range resp {
		switch r.Status {
		case "rejected":
			return fmt.Errorf("Email to %s (id: %s) rejected: %v\n", r.Email,
				r.Id, r.RejectedReason)
		case "invalid":
			return fmt.Errorf("Error sending invalid message %s to %s\n", r.Id,
				r.Email)
		case "queued", "sent":
			fmt.Printf("Email to %s (id: %s) %s\n", r.Email, r.Id, r.Status)
		default:
			fmt.Printf("Impossible status `%s` seen for SendResponse `%#v`\n",
				r.Status, r)
		}
	}
	return nil
}
