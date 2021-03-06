package agent_docker_compose

import (
	"errors"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type DockerCompose struct {
	Version  string   `yaml:"version"`
	Services Services `yaml:"services"`
	Networks Networks `yaml:"networks"`
}
type Build struct {
	Context    string `yaml:"context"`
	Dockerfile string `yaml:"dockerfile"`
}

type Services struct {
	Issuer1   Issuer   `yaml:"issuer1,omitempty"`
	Issuer2   Issuer   `yaml:"issuer2,omitempty"`
	Issuer3   Issuer   `yaml:"issuer3,omitempty"`
	Issuer4   Issuer   `yaml:"issuer4,omitempty"`
	Issuer5   Issuer   `yaml:"issuer5,omitempty"`
	Holder1   Holder   `yaml:"holder1,omitempty"`
	Holder2   Holder   `yaml:"holder2,omitempty"`
	Holder3   Holder   `yaml:"holder3,omitempty"`
	Holder4   Holder   `yaml:"holder4,omitempty"`
	Holder5   Holder   `yaml:"holder5,omitempty"`
	Verifier1 Verifier `yaml:"verifier1,omitempty"`
	Verifier2 Verifier `yaml:"verifier2,omitempty"`
	Verifier3 Verifier `yaml:"verifier3,omitempty"`
	Verifier4 Verifier `yaml:"verifier4,omitempty"`
	Verifier5 Verifier `yaml:"verifier5,omitempty"`
}
type External struct {
	Name string `yaml:"name"`
}
type Default struct {
	External External `yaml:"external"`
}
type Networks struct {
	Default Default `yaml:"default"`
}

type Issuer struct {
	Build   Build    `yaml:"build"`
	Ports   []string `yaml:"ports"`
	Command string   `yaml:"command"`
	Volumes []string `yaml:"volumes"`
}

type Holder struct {
	Build   Build    `yaml:"build"`
	Ports   []string `yaml:"ports"`
	Command string   `yaml:"command"`
	Volumes []string `yaml:"volumes"`
}

type Verifier struct {
	Build   Build    `yaml:"build"`
	Ports   []string `yaml:"ports"`
	Command string   `yaml:"command"`
	Volumes []string `yaml:"volumes"`
}

// dot -> graph -> docker-compose.yml
func ConvertFromGraph(dotPath string, workdir string, networkName string, myIPAddress string, agentNameAndSeed map[string]string, ngrokUrlList map[string]string) error {
	bytes, err := ioutil.ReadFile(dotPath)
	if err != nil {
		return errors.New("can't read dot file")
	}
	graph, err := gographviz.Read(bytes)
	if err != nil {
		return errors.New("can't parse dot file")
	}

	d := DockerCompose{}
	d.Version = "3"
	d.Networks.Default.External.Name = networkName

	//Agent Counter
	issuerNum := 0
	holderNum := 0
	verifierNum := 0

	for _, node := range graph.Nodes.Nodes {
		//fmt.Println(node)
		label, _ := attrToBetter(node.Attrs["label"], node.Attrs["xlabel"])
		seed := agentNameAndSeed[node.Name]
		//fmt.Println(attrs)
		switch label {
		case "Issuer":
			issuerNum += 1
			//fmt.Printf("%v is Issuer\n", node.Name)

			// ~~TODO interfaceとか使って書く~~
			// TODO yamlのunmarshalする構造体を配列を使って書いて、頭の悪さを払拭する
			// TODO :2021/03/06 :SSHポートフォワーディング等を使用して、無限にURLを生成する
			// NOTE : ngrokの制約により、Issuer x1(:8001) / Holder x1(:8004) / Verifier x1(:8007)
			switch issuerNum {
			case 1:
				d.Services.Issuer1.Build.Context = "./aries-cloudagent-python"
				d.Services.Issuer1.Build.Dockerfile = "./docker/Dockerfile.run"
				d.Services.Issuer1.Ports = []string{"8001:8000", "11000-11999:11000"}
				d.Services.Issuer1.Command = getAgentCommand(node.Name, myIPAddress, seed, ngrokUrlList["issuer1"])
				d.Services.Issuer1.Volumes = []string{"./aries-cloudagent-python/logs/:/home/indy/logs"}
				//case 2:
				//	d.Services.Issuer2.Build.Context = "./aries-cloudagent-python"
				//	d.Services.Issuer2.Build.Dockerfile = "./docker/Dockerfile.run"
				//	d.Services.Issuer2.Ports = []string{"8002:8000", "11000-11999:11000"}
				//	d.Services.Issuer2.Command = getAgentCommand(node.Name, myIPAddress, seed, ngrokUrlList["issuer2"])
				//	d.Services.Issuer2.Volumes = []string{"./aries-cloudagent-python/logs/:/home/indy/logs"}
				//case 3:
				//	d.Services.Issuer3.Build.Context = "./aries-cloudagent-python"
				//	d.Services.Issuer3.Build.Dockerfile = "./docker/Dockerfile.run"
				//	d.Services.Issuer3.Ports = []string{"8003:8000", "11000-11999:11000"}
				//	d.Services.Issuer3.Command = getAgentCommand(node.Name, myIPAddress, seed)
				//	d.Services.Issuer3.Volumes = []string{"./aries-cloudagent-python/logs/:/home/indy/logs"}
			}

		case "Holder":
			holderNum += 1
			switch holderNum {
			case 1:
				d.Services.Holder1.Build.Context = "./aries-cloudagent-python"
				d.Services.Holder1.Build.Dockerfile = "./docker/Dockerfile.run"
				d.Services.Holder1.Ports = []string{"8004:8000", "11000-11999:11000"}
				d.Services.Holder1.Command = getAgentCommand(node.Name, myIPAddress, seed, ngrokUrlList["holder1"])
				d.Services.Holder1.Volumes = []string{"./aries-cloudagent-python/logs/:/home/indy/logs"}
				//case 2:
				//	d.Services.Holder2.Build.Context = "./aries-cloudagent-python"
				//	d.Services.Holder2.Build.Dockerfile = "./docker/Dockerfile.run"
				//	d.Services.Holder2.Ports = []string{"8005:8000", "11000-11999:11000"}
				//	d.Services.Holder2.Command = getAgentCommand(node.Name, myIPAddress, seed, ngrokUrlList["holder2"])
				//	d.Services.Holder2.Volumes = []string{"./aries-cloudagent-python/logs/:/home/indy/logs"}
				//case 3:
				//	d.Services.Holder3.Build.Context = "./aries-cloudagent-python"
				//	d.Services.Holder3.Build.Dockerfile = "./docker/Dockerfile.run"
				//	d.Services.Holder3.Ports = []string{"8006:8000", "11000-11999:11000"}
				//	d.Services.Holder3.Command = getAgentCommand(node.Name, myIPAddress, seed)
				//	d.Services.Holder3.Volumes = []string{"./aries-cloudagent-python/logs/:/home/indy/logs"}
			}
			//fmt.Printf("%v is Holder\n", node.Name)

		case "Verifier":
			verifierNum += 1
			switch verifierNum {
			case 1:
				d.Services.Verifier1.Build.Context = "./aries-cloudagent-python"
				d.Services.Verifier1.Build.Dockerfile = "./docker/Dockerfile.run"
				d.Services.Verifier1.Ports = []string{"8007:8000", "11000-11999:11000"}
				d.Services.Verifier1.Command = getAgentCommand(node.Name, myIPAddress, seed, ngrokUrlList["verifier1"])
				d.Services.Verifier1.Volumes = []string{"./aries-cloudagent-python/logs/:/home/indy/logs"}
			}

		}
	}
	out, err := yaml.Marshal(&d)
	if err != nil {
		return errors.New("can't serialize data")
	}
	//fmt.Println(out)

	dstPath := filepath.Join(workdir, "docker-compose.yml")
	err = ioutil.WriteFile(dstPath, out, 0644)
	if err != nil {
		return errors.New("can't write to yml file")
	}
	return nil
}

func attrToBetter(label string, xlabel string) (string, map[string]string) {
	label = strings.Trim(label, "\"")

	//TODO xlabelをもとに属性返す
	if xlabel == "" {
		return label, map[string]string{}
	}

	trimmed := strings.Trim(xlabel, "\"")
	splitted := strings.Split(trimmed, ",")
	attrs := map[string]string{}
	for _, chunk := range splitted {
		div := strings.Split(chunk, ":")
		attr := div[0]
		value := div[1]
		attrs[attr] = value
	}
	return label, attrs
}

func getAgentCommand(label string, ip string, seed string, ngrokUrl string) string {
	cmd := fmt.Sprintf(
		"start --label %v --inbound-transport http 0.0.0.0 8000 --outbound-transport http --admin 0.0.0.0 11000 --admin-insecure-mode --genesis-url http://%v:9000/genesis --seed %v --wallet-type indy --wallet-name %v --wallet-key welldone --endpoint %v --public-invites --auto-accept-invites --auto-accept-requests --auto-ping-connection --debug-connections", label, ip, seed, label, ngrokUrl)
	return cmd
}

func GetServiceNameFromWorkdir(workdir string) ([]string, error) {
	ymlPath := filepath.Join(workdir, "docker-compose.yml")
	bytes, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		return nil, err
	}

	d := DockerCompose{}
	err = yaml.Unmarshal(bytes, &d)
	if err != nil {
		return nil, err
	}

	return nil, nil

	//for _, service := range d.Services
}

func simpleRead(yamlPath string) error {
	bytes, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return errors.New("can't open yaml file")
	}

	d := DockerCompose{}
	err = yaml.Unmarshal(bytes, &d)
	if err != nil {
		return errors.New("can't parse yml file")
	}

	fmt.Println(d)
	return nil
}
