package issuer

type Issuer struct {
	Schema []Schema `yaml:"schema"`
}
type Schema struct {
	Name    string   `yaml:"name"`
	Version string   `yaml:"version"`
	Attr    []string `yaml:"attr"`
}

func CreateIssuerConf(issuerName string, issuerAttribute []string) error {
	issuerConf := Issuer{}
	sc := Schema{}
	sc.Name = issuerName
	for _, attr := range issuerAttribute{
		sc.Attr = append(sc.Attr, attr)
	}
}
