package commander

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type PublicDid struct {
	Did string
	Seed string
	Verkey string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RegisterDID(agentNameList []string) (map[string]string, error){
	agentAndSeed, err := getSeedList(agentNameList)
	if err != nil {
		return nil, err
	}

	agentAndDid := map[string]string{}
	for agent, seed := range agentAndSeed{
		publicDid, err := ComLedger(agent, seed)
		if err != nil{
			return nil, err
		}
		agentAndDid[agent] = publicDid
	}

	return agentAndDid, nil
}

func ComLedger(alias string, seed string) (string, error){
	url := "http://localhost:9000/register"

	jsonData := `{"alias":"` + alias + `","seed":"` + seed + `","role":"TRUST_ANCHOR"}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonData)),
	)
	if err != nil{
		return "", err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var d PublicDid
	err = json.Unmarshal(body, &d)
	if err != nil{
		return "", err
	}

	return d.Did, nil
}

func getSeedList(agentNameList []string) (map[string]string, error){
	seedList := map[string]string{}
	for _, name := range agentNameList{
		seedList[name] = getRandomString(32)
	}
	return seedList, nil
}

func getRandomString(n int) string{
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
