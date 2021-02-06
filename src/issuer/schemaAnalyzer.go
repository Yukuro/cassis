package issuer

import "encoding/json"

type SchemaResponse struct {
	SchemaID string `json:"schema_id"`
	Schema   Schema `json:"schema"`
}
type Schema struct {
	Ver       string   `json:"ver"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	AttrNames []string `json:"attrNames"`
	SeqNo     int      `json:"seqNo"`
}

func ExtractSchemaNameAndSchemaId(body []byte) (string, string, error) {
	sr := SchemaResponse{}

	err := json.Unmarshal(body, &sr)
	if err != nil {
		return "", "", err
	}

	return sr.Schema.Name, sr.SchemaID, nil
}
