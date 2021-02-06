package commander

import (
	"bytes"
	"cli/src/issuer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type PublicDid struct {
	Did    string
	Seed   string
	Verkey string
}

type CreateInvitation struct {
	ConnectionID  string     `json:"connection_id"`
	Invitation    Invitation `json:"invitation"`
	InvitationURL string     `json:"invitation_url"`
	Alias         string     `json:"alias"`
}
type Invitation struct {
	Type            string   `json:"@type"`
	ID              string   `json:"@id"`
	Label           string   `json:"label"`
	RecipientKeys   []string `json:"recipientKeys"`
	ServiceEndpoint string   `json:"serviceEndpoint"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RegisterDID(agentNameList []string) (map[string]string, error) {
	agentAndSeed, err := getSeedList(agentNameList)
	if err != nil {
		return nil, err
	}
	//dbp(agentAndSeed)

	agentAndDid := map[string]string{}
	for agent, seed := range agentAndSeed {
		publicDid, err := ComLedger(agent, seed)
		fmt.Printf("Registering %v(seed: %v ) to ledger\n", agent, seed)
		time.Sleep(time.Second * 5)
		if err != nil {
			//return nil, err
			panic(err)
		}
		agentAndDid[agent] = publicDid
	}
	//fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAa")
	//dbp(agentAndSeed)
	return agentAndSeed, nil
}

func ComLedger(alias string, seed string) (string, error) {
	url := "http://localhost:9000/register"

	jsonData := `{"alias":"` + alias + `","seed":"` + seed + `","role":"TRUST_ANCHOR"}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonData)),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var d PublicDid
	err = json.Unmarshal(body, &d)
	if err != nil {
		return "", err
	}

	return d.Did, nil
}

func InvitationToHolder(targetUrl string, agentName string) (string, string, error) {
	targetUrl = targetUrl + "/connections/create-invitation"

	alias := fmt.Sprintf("%v_%v", agentName, GetRandomString(5))

	jsonData := `{"alias":"` + alias + `","auto_accept":true}`

	req, err := http.NewRequest(
		"POST",
		targetUrl,
		bytes.NewBuffer([]byte(jsonData)),
	)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("{\"alias\": \"%v\", \"auto_accept\": true} --> %v\n", alias, targetUrl)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var d CreateInvitation
	err = json.Unmarshal(body, &d)
	if err != nil {
		return "", "", nil
	}

	connectionId := d.ConnectionID

	bytes, err := json.Marshal(&d.Invitation)
	if err != nil {
		return "", "", nil
	}

	return connectionId, string(bytes), nil
}

func OriginateSchema(issuerUrl string, schemaName string, schemaVersion string, schemaAttr []string) (string, string, []string, string, error) {
	jsonData, err := issuer.PackSchema(schemaName, schemaVersion, schemaAttr)
	if err != nil {
		return "", "", []string{}, "", err
	}

	req, err := http.NewRequest(
		"POST",
		issuerUrl,
		bytes.NewBuffer([]byte(jsonData)),
	)
	if err != nil {
		return "", "", []string{}, "", err
	}

	//fmt.Printf("POST\n%v\n", string(jsonData))

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", []string{}, "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", []string{}, "", err
	}
	originatedSchemaName, originatedSchemaId, err := issuer.ExtractSchemaNameAndSchemaId(body)
	if err != nil {
		return "", "", []string{}, "", err
	}

	return originatedSchemaName, schemaVersion, schemaAttr, originatedSchemaId, nil
}

func getSeedList(agentNameList []string) (map[string]string, error) {
	seedList := map[string]string{}
	for _, name := range agentNameList {
		seedList[name] = GetRandomString(32)
	}
	return seedList, nil
}

func GetRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func dbp(seed map[string]string) {
	if seed != nil {
		fmt.Println(seed)
	} else {
		panic(seed)
	}
}
