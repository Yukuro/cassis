package post_issuecredential_records_credexid_sendrequest

import "encoding/json"

type Response struct {
	ThreadID               string `json:"thread_id"`
	AutoIssue              bool   `json:"auto_issue"`
	SchemaID               string `json:"schema_id"`
	UpdatedAt              string `json:"updated_at"`
	CredentialDefinitionID string `json:"credential_definition_id"`
	CreatedAt              string `json:"created_at"`
	Trace                  bool   `json:"trace"`
	ConnectionID           string `json:"connection_id"`
	CredentialExchangeID   string `json:"credential_exchange_id"`
	CredentialProposalDict struct {
		Type               string `json:"@type"`
		ID                 string `json:"@id"`
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
		CredDefID string `json:"cred_def_id"`
	} `json:"credential_proposal_dict"`
	AutoOffer       bool `json:"auto_offer"`
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
	CredentialRequestMetadata struct {
		MasterSecretBlindingData struct {
			VPrime  string      `json:"v_prime"`
			VrPrime interface{} `json:"vr_prime"`
		} `json:"master_secret_blinding_data"`
		Nonce            string `json:"nonce"`
		MasterSecretName string `json:"master_secret_name"`
	} `json:"credential_request_metadata"`
	Role              string `json:"role"`
	AutoRemove        bool   `json:"auto_remove"`
	State             string `json:"state"`
	Initiator         string `json:"initiator"`
	CredentialRequest struct {
		ProverDid string `json:"prover_did"`
		CredDefID string `json:"cred_def_id"`
		BlindedMs struct {
			U                   string      `json:"u"`
			Ur                  interface{} `json:"ur"`
			HiddenAttributes    []string    `json:"hidden_attributes"`
			CommittedAttributes struct {
			} `json:"committed_attributes"`
		} `json:"blinded_ms"`
		BlindedMsCorrectnessProof struct {
			C        string `json:"c"`
			VDashCap string `json:"v_dash_cap"`
			MCaps    struct {
				MasterSecret string `json:"master_secret"`
			} `json:"m_caps"`
			RCaps struct {
			} `json:"r_caps"`
		} `json:"blinded_ms_correctness_proof"`
		Nonce string `json:"nonce"`
	} `json:"credential_request"`
}

func GetStateFromBody(body []byte) (string, error) {
	resp := Response{}

	err := json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}

	return resp.State, nil
}
