package get_wallet_did

import "encoding/json"

type WalletDid struct {
	Results []struct {
		Did     string `json:"did"`
		Verkey  string `json:"verkey"`
		Posture string `json:"posture"`
	} `json:"results"`
}

func AnalyzeWalletDid(body []byte) (string, error) {
	wd := WalletDid{}

	err := json.Unmarshal(body, &wd)
	if err != nil {
		return "", err
	}

	return wd.Results[0].Did, nil
}
