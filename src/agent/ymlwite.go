package agent

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Conf struct {
	Issuer Issuer      `yaml:"issuer"`
	Ledger interface{} `yaml:"ledger"`
}
type Schema struct {
	Name    string   `yaml:"name"`
	Version string   `yaml:"version"`
	Attr    []string `yaml:"attr"`
	ID      string   `yaml:"id,omitempty"`
}
type Issuer struct {
	Schema []Schema `yaml:"schema"`
}

func CreateNilConf(dstPath string) error {
	cf := Conf{}
	is := Issuer{}
	//is.Agent = []string{}

	sc := Schema{
		Name:    "",
		Version: "",
		Attr:    []string{},
	}
	is.Schema = append(is.Schema, sc)
	cf.Issuer = is

	out, err := yaml.Marshal(&cf)
	if err != nil {
		return err
	}

	dstPath = filepath.Join(dstPath, "config.yml")
	err = ioutil.WriteFile(dstPath, out, 0666)
	if err != nil {
		return err
	}
	return nil
}

func CreateIssuerConf(dstPath string, issuerName string, issuerVersion string, issuerAttribute []string, schemaId string) error {
	cf := Conf{}
	is := Issuer{}
	//is.Agent = []string{}

	sc := Schema{
		Name:    issuerName,
		Version: issuerVersion,
		Attr:    issuerAttribute,
		ID:      schemaId,
	}
	is.Schema = append(is.Schema, sc)
	cf.Issuer = is

	out, err := yaml.Marshal(&cf)
	if err != nil {
		return err
	}

	dstPath = filepath.Join(dstPath, "config.yml")
	err = ioutil.WriteFile(dstPath, out, 0666)
	if err != nil {
		return err
	}

	return nil
}

func AnalyzeIssuerConf(initDir string) (map[string][]string, error) {
	ymlPath := filepath.Join(initDir, "config.yml")
	bytes, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		return nil, err
	}

	cf := Conf{}
	err = yaml.Unmarshal(bytes, &cf)
	if err != nil {
		return nil, err
	}

	res := make(map[string][]string)
	for _, sc := range cf.Issuer.Schema {
		// TODO Version 1.0固定をやめる
		res[sc.Name] = sc.Attr
	}

	return res, nil
}
