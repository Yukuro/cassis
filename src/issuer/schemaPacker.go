package issuer

import "encoding/json"

type SchemaRequest struct {
	Attributes    []string `json:"attributes"`
	SchemaName    string   `json:"schema_name"`
	SchemaVersion string   `json:"schema_version"`
}

func PackSchema(schemaName string, schemaVersion string, schemaAttr []string) ([]byte, error) {
	sc := SchemaRequest{}
	sc.Attributes = schemaAttr
	sc.SchemaName = schemaName
	sc.SchemaVersion = schemaVersion

	out, err := json.Marshal(sc)
	if err != nil {
		return nil, err
	}
	return out, nil
}
