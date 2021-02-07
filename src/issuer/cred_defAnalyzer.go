package issuer

import "encoding/json"

type Cred_defResponse struct {
	CredentialDefinitionID string `json:"credential_definition_id"`
}

func AnalyzeCred_def(body []byte) (string, error) {
	cr := Cred_defResponse{}

	err := json.Unmarshal(body, &cr)
	if err != nil {
		return "", err
	}

	return cr.CredentialDefinitionID, nil
}
