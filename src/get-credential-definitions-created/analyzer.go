package get_credential_definitions_created

import "encoding/json"

type CreatedCredDef struct {
	CredentialDefinitionIds []string `json:"credential_definition_ids"`
}

func AnalyzeCreatedCredDef(body []byte) ([]string, error) {
	cd := CreatedCredDef{}

	err := json.Unmarshal(body, &cd)
	if err != nil {
		return []string{}, err
	}

	return cd.CredentialDefinitionIds, nil
}
