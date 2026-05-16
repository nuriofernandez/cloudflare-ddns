package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var token = os.Getenv("IPINFO_API_TOKEN")

type IPInfoResponse struct {
	Ip            string `json:"ip"`
	Asn           string `json:"asn"`
	AsName        string `json:"as_name"`
	AsDomain      string `json:"as_domain"`
	CountryCode   string `json:"country_code"`
	Country       string `json:"country"`
	ContinentCode string `json:"continent_code"`
	Continent     string `json:"continent"`
}

func FetchIp() (string, error) {
	// Perform the API call
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.ipinfo.io/lite/me?token="+token, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Parse response
	defer resp.Body.Close()
	response := IPInfoResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	// Log and return!
	fmt.Printf("[IPInfo] Fetched current IP. (%s@%s)\n", response.Asn, response.Ip)
	return response.Ip, nil
}
