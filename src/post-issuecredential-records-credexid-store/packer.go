package post_issuecredential_records_credexid_store

import "encoding/json"

type Records struct {
	CredentialID string `json:"credential_id"`
}

func PackStoreCredential() ([]byte, error) {
	rc := Records{}

	rc.CredentialID = "string"

	out, err := json.Marshal(rc)
	if err != nil {
		return nil, err
	}

	return out, nil
}
