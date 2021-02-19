package ngrok_analyze

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Inspect struct {
	Tunnels []struct {
		Name      string `json:"name"`
		URI       string `json:"uri"`
		PublicURL string `json:"public_url"`
		Proto     string `json:"proto"`
		Config    struct {
			Addr    string `json:"addr"`
			Inspect bool   `json:"inspect"`
		} `json:"config"`
		Metrics struct {
			Conns struct {
				Count  int `json:"count"`
				Gauge  int `json:"gauge"`
				Rate1  int `json:"rate1"`
				Rate5  int `json:"rate5"`
				Rate15 int `json:"rate15"`
				P50    int `json:"p50"`
				P90    int `json:"p90"`
				P95    int `json:"p95"`
				P99    int `json:"p99"`
			} `json:"conns"`
			HTTP struct {
				Count  int `json:"count"`
				Rate1  int `json:"rate1"`
				Rate5  int `json:"rate5"`
				Rate15 int `json:"rate15"`
				P50    int `json:"p50"`
				P90    int `json:"p90"`
				P95    int `json:"p95"`
				P99    int `json:"p99"`
			} `json:"http"`
		} `json:"metrics"`
	} `json:"tunnels"`
	URI string `json:"uri"`
}

func GetNgrokUrl() (map[string]string, error) {
	url := "http://localhost:4040/api/tunnels"

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var inspect Inspect
	err = json.Unmarshal(body, &inspect)
	if err != nil {
		return nil, err
	}

	// TODO URLを返す
	fmt.Println("All Go!")

	urlList := map[string]string{}
	for _, tunnel := range inspect.Tunnels {
		urlList[tunnel.Name] = tunnel.PublicURL
	}

	return urlList, nil
}
