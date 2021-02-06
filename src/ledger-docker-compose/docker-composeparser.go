package ledger_docker_compose

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type DockerCompose struct {
	Version  string   `yaml:"version"`
	Services Services `yaml:"services"`
	Networks Networks `yaml:"networks"`
	Volumes  Volumes  `yaml:"volumes"`
}
type Client struct {
	Image       string   `yaml:"image"`
	Command     string   `yaml:"command"`
	Environment []string `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
}
type Webserver struct {
	Image       string   `yaml:"image"`
	Command     string   `yaml:"command"`
	Environment []string `yaml:"environment"`
	Ports       []string `yaml:"ports"`
	Volumes     []string `yaml:"volumes"`
}
type Synctest struct {
	Image       string   `yaml:"image"`
	Command     string   `yaml:"command"`
	Environment []string `yaml:"environment"`
	Ports       []string `yaml:"ports"`
	Volumes     []string `yaml:"volumes"`
}
type Nodes struct {
	Image       string   `yaml:"image"`
	Command     string   `yaml:"command"`
	Ports       []string `yaml:"ports"`
	Environment []string `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
}
type Node1 struct {
	Image       string   `yaml:"image"`
	Command     string   `yaml:"command"`
	Ports       []string `yaml:"ports"`
	Environment []string `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
}
type Node2 struct {
	Image       string   `yaml:"image"`
	Command     string   `yaml:"command"`
	Ports       []string `yaml:"ports"`
	Environment []string `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
}
type Node3 struct {
	Image       string   `yaml:"image"`
	Command     string   `yaml:"command"`
	Ports       []string `yaml:"ports"`
	Environment []string `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
}
type Node4 struct {
	Image       string   `yaml:"image"`
	Command     string   `yaml:"command"`
	Ports       []string `yaml:"ports"`
	Environment []string `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
}
type Services struct {
	Client    Client    `yaml:"client"`
	Webserver Webserver `yaml:"webserver"`
	Synctest  Synctest  `yaml:"synctest"`
	Nodes     Nodes     `yaml:"nodes"`
	Node1     Node1     `yaml:"node1"`
	Node2     Node2     `yaml:"node2"`
	Node3     Node3     `yaml:"node3"`
	Node4     Node4     `yaml:"node4"`
}
type External struct {
	Name interface{} `yaml:"name"`
}
type Default struct {
	External External `yaml:"external"`
}
type Networks struct {
	Default Default `yaml:"default"`
}
type Volumes struct {
	ClientData      interface{} `yaml:"client-data"`
	WebserverCli    interface{} `yaml:"webserver-cli"`
	WebserverLedger interface{} `yaml:"webserver-ledger"`
	Node1Data       interface{} `yaml:"node1-data"`
	Node2Data       interface{} `yaml:"node2-data"`
	Node3Data       interface{} `yaml:"node3-data"`
	Node4Data       interface{} `yaml:"node4-data"`
	NodesData       interface{} `yaml:"nodes-data"`
}

func RenameNetworks(workdir string, networkName string) error {
	yamlPath := filepath.Join(workdir, "von-network", "docker-compose.yml")
	bytes, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return errors.New("can't open yaml file")
	}

	d := DockerCompose{}
	err = yaml.Unmarshal(bytes, &d)
	if err != nil {
		return errors.New("can't parse yml file")
	}

	d.Networks.Default.External.Name = networkName

	out, err := yaml.Marshal(&d)
	if err != nil {
		errors.New("can't serialize data")
	}

	dstPath := filepath.Join(workdir, "von-network", "docker-compose.yml")
	err = ioutil.WriteFile(dstPath, out, 0644)
	if err != nil {
		errors.New("can't write to yml file")
	}
	//fmt.Println(out)
	return nil
}
