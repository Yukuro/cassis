package post_issuecredential_send

import "encoding/json"

type IssueCredential struct {
	AutoRemove         bool               `json:"auto_remove,omitempty"`
	Comment            string             `json:"comment,omitempty"`
	ConnectionID       string             `json:"connection_id"`
	CredDefID          string             `json:"cred_def_id,omitempty"`
	CredentialProposal CredentialProposal `json:"credential_proposal"`
	IssuerDid          string             `json:"issuer_did,omitempty"`
	SchemaID           string             `json:"schema_id,omitempty"`
	SchemaIssuerDid    string             `json:"schema_issuer_did,omitempty"`
	SchemaName         string             `json:"schema_name,omitempty"`
	SchemaVersion      string             `json:"schema_version,omitempty"`
	Trace              bool               `json:"trace,omitempty"`
}
type Attributes struct {
	MimeType string `json:"mime-type"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}
type CredentialProposal struct {
	Type       string       `json:"@type"`
	Attributes []Attributes `json:"attributes"`
}

func PackIssueCredential(connectionId string, attributes map[string]string, issuerDid string, schemaId string, credDefId string) ([]byte, error) {
	ic := IssueCredential{}

	ic.ConnectionID = connectionId
	ic.CredDefID = credDefId
	ic.CredentialProposal.Type = "issue-credential/1.0/credential-preview"

	for attr, value := range attributes {
		at := Attributes{}
		at.MimeType = "text/plain"
		at.Name = attr
		at.Value = value

		ic.CredentialProposal.Attributes = append(ic.CredentialProposal.Attributes, at)
	}

	ic.IssuerDid = issuerDid
	ic.SchemaID = schemaId

	out, err := json.Marshal(ic)
	if err != nil {
		return nil, err
	}

	return out, nil
}
