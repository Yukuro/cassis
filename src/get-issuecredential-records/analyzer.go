package get_issuecredential_records

import "encoding/json"

type Records struct {
	Results []struct {
		CredentialDefinitionID string `json:"credential_definition_id"`
		CreatedAt              string `json:"created_at"`
		Initiator              string `json:"initiator"`
		AutoIssue              bool   `json:"auto_issue"`
		UpdatedAt              string `json:"updated_at"`
		ThreadID               string `json:"thread_id"`
		AutoOffer              bool   `json:"auto_offer"`
		CredentialProposalDict struct {
			Type               string `json:"@type"`
			ID                 string `json:"@id"`
			CredDefID          string `json:"cred_def_id"`
			Comment            string `json:"comment"`
			SchemaID           string `json:"schema_id"`
			CredentialProposal struct {
				Type       string `json:"@type"`
				Attributes []struct {
					Name     string `json:"name"`
					MimeType string `json:"mime-type"`
					Value    string `json:"value"`
				} `json:"attributes"`
			} `json:"credential_proposal"`
		} `json:"credential_proposal_dict"`
		ConnectionID    string `json:"connection_id"`
		Role            string `json:"role"`
		CredentialOffer struct {
			SchemaID            string `json:"schema_id"`
			CredDefID           string `json:"cred_def_id"`
			KeyCorrectnessProof struct {
				C     string     `json:"c"`
				XzCap string     `json:"xz_cap"`
				XrCap [][]string `json:"xr_cap"`
			} `json:"key_correctness_proof"`
			Nonce string `json:"nonce"`
		} `json:"credential_offer"`
		SchemaID             string `json:"schema_id"`
		Trace                bool   `json:"trace"`
		CredentialExchangeID string `json:"credential_exchange_id"`
		AutoRemove           bool   `json:"auto_remove"`
		State                string `json:"state"`
	} `json:"results"`
}

func GetOfferCredExIdListFromBody(body []byte) ([]string, error) {
	rec := Records{}

	err := json.Unmarshal(body, &rec)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, record := range rec.Results {
		if record.State == "offer_received" {
			out = append(out, record.CredentialExchangeID)
		}
	}

	return out, nil
}
