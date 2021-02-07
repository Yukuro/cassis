package issuer

import "encoding/json"

type Cred_defRequest struct {
	RevocationRegistrySize int    `json:"revocation_registry_size"`
	SchemaID               string `json:"schema_id"`
	SupportRevocation      bool   `json:"support_revocation"`
	Tag                    string `json:"tag"`
}

func PackCred_def(schemaId string) ([]byte, error) {
	cr := Cred_defRequest{}
	cr.RevocationRegistrySize = 1000
	cr.SchemaID = schemaId
	cr.SupportRevocation = false
	cr.Tag = "default"

	out, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}
	return out, nil
}
