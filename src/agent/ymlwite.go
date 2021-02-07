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
	Name       string   `yaml:"name"`
	Version    string   `yaml:"version"`
	Attr       []string `yaml:"attr"`
	ID         string   `yaml:"id,omitempty"`
	Cred_defId string   `yaml:"cred_defid,omitempty"`
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

func CreateIssuerConf(dstPath string, schemaName string, schemaVersion string, schemaAttributes []string, schemaId string, cred_defId string) error {
	cf := Conf{}
	is := Issuer{}
	//is.Agent = []string{}

	sc := Schema{
		Name:       schemaName,
		Version:    schemaVersion,
		Attr:       schemaAttributes,
		ID:         schemaId,
		Cred_defId: cred_defId,
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

func AddIssuerConfWithWorkdir(dstPath string, schemaName string, schemaVersion string, schemaAttributes []string, schemaId string, cred_defId string) error {
	ymlPath := filepath.Join(".cassis", "config.yml")
	bytes, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		return err
	}

	cf := Conf{}
	err = yaml.Unmarshal(bytes, &cf)
	if err != nil {
		return err
	}

	//TODO 全属性に対応させる
	if cred_defId != "" {
		cf.Issuer.Schema[0].Cred_defId = cred_defId
	}

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

func AnalyzeIssuerConf(initDir string) ([]map[string]string, error) {
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

	var res []map[string]string
	for _, sc := range cf.Issuer.Schema {
		// TODO Version 1.0固定をやめる
		tp := map[string]string{}
		tp["name"] = sc.Name
		tp["version"] = sc.Version
		tp["id"] = sc.ID
		res = append(res, tp)
	}

	return res, nil
}
