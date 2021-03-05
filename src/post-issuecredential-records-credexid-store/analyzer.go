package post_issuecredential_records_credexid_store

import "encoding/json"

type Response struct {
	AutoIssue    bool   `json:"auto_issue"`
	AutoOffer    bool   `json:"auto_offer"`
	AutoRemove   bool   `json:"auto_remove"`
	ConnectionID string `json:"connection_id"`
	CreatedAt    string `json:"created_at"`
	Credential   struct {
	} `json:"credential"`
	CredentialDefinitionID string `json:"credential_definition_id"`
	CredentialExchangeID   string `json:"credential_exchange_id"`
	CredentialID           string `json:"credential_id"`
	CredentialOffer        struct {
	} `json:"credential_offer"`
	CredentialOfferDict struct {
	} `json:"credential_offer_dict"`
	CredentialProposalDict struct {
	} `json:"credential_proposal_dict"`
	CredentialRequest struct {
	} `json:"credential_request"`
	CredentialRequestMetadata struct {
	} `json:"credential_request_metadata"`
	ErrorMsg       string `json:"error_msg"`
	Initiator      string `json:"initiator"`
	ParentThreadID string `json:"parent_thread_id"`
	RawCredential  struct {
	} `json:"raw_credential"`
	RevocRegID   string `json:"revoc_reg_id"`
	RevocationID string `json:"revocation_id"`
	Role         string `json:"role"`
	SchemaID     string `json:"schema_id"`
	State        string `json:"state"`
	ThreadID     string `json:"thread_id"`
	Trace        bool   `json:"trace"`
	UpdatedAt    string `json:"updated_at"`
}

func GetCredentialIdFromBody(body []byte) (string, error) {
	resp := Response{}

	err := json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}

	return resp.CredentialID, nil
}
