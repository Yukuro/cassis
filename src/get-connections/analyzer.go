package get_connections

import "encoding/json"

type Response struct {
	Results []struct {
		UpdatedAt      string `json:"updated_at"`
		TheirLabel     string `json:"their_label,omitempty"`
		InvitationKey  string `json:"invitation_key"`
		Accept         string `json:"accept"`
		TheirDid       string `json:"their_did,omitempty"`
		MyDid          string `json:"my_did,omitempty"`
		State          string `json:"state"`
		InvitationMode string `json:"invitation_mode"`
		ConnectionID   string `json:"connection_id"`
		Initiator      string `json:"initiator"`
		CreatedAt      string `json:"created_at"`
		RoutingState   string `json:"routing_state"`
	} `json:"results"`
}

func GetActiveConnections(body []byte) ([]string, error) {
	resp := Response{}

	err := json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, res := range resp.Results {
		if res.State == "active" {
			out = append(out, res.ConnectionID)
		}
	}

	return out, nil
}
