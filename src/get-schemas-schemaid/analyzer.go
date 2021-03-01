package get_schemas_schemaid

import (
	"encoding/json"
)

type SchemaFromId struct {
	Schema struct {
		AttrNames []string `json:"attrNames"`
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		SeqNo     int      `json:"seqNo"`
		Ver       string   `json:"ver"`
		Version   string   `json:"version"`
	} `json:"schema"`
}

func GetAttrNamesFromBody(body []byte) ([]string, error) {
	sf := SchemaFromId{}

	err := json.Unmarshal(body, &sf)
	if err != nil {
		return nil, err
	}

	return sf.Schema.AttrNames, nil
}
