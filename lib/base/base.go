package base

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/oauth2"
)

type Base struct {
	*http.Client
	apiKey    string
	apiSecret string
	baseUrl   string
	debug     bool
}

/// Create new base instance with the provided credentials
func NewBase(apiKey string, apiSecret string, useSandbox bool, debug bool) Base {
	// Bind API keys
	a := Base{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseUrl:   "https://api.business.govt.nz/",
		debug:     debug,
	}

	if useSandbox {
		a.baseUrl = "https://sandbox.api.business.govt.nz/"
	}

	// Setup oauth client
	a.Client = oauth2.NewClient(context.Background(), &a)

	return a
}

/// Create new base instance from environment flags
func NewBaseFromEnv() (*Base, error) {
	// Fetch key and secret from environment
	apiKey := os.Getenv("BUSINESSNZ_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("no BUSINESSNZ_API_KEY variable found")
	}

	apiSecret := os.Getenv("BUSINESSNZ_API_SECRET")
	if apiSecret == "" {
		return nil, fmt.Errorf("no BUSINESSNZ_API_SECRET variable found")
	}

	// Override base with sandbox URL if flag is set
	useSandbox := (os.Getenv("BUSINESSNZ_API_SANDBOX") != "")

	debug := (os.Getenv("BUSINESSNZ_API_DEBUG") != "")

	// Create new API instance
	base := NewBase(apiKey, apiSecret, useSandbox, debug)

	return &base, nil
}

// `oauth2.TokenSource` implementation to fetch BusinessNZ API tokens
func (a *Base) Token() (*oauth2.Token, error) {

	v := url.Values{}
	v.Add("grant_type", "client_credentials")

	// Setup POST to request new token
	// Note this address is shared for the sandbox and production interfaces
	req, err := http.NewRequest("POST", "https://api.business.govt.nz/services/token?"+v.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request object: %v", err)
	}

	// Encode API key and secret
	auth := base64.StdEncoding.EncodeToString([]byte(a.apiKey + ":" + a.apiSecret))
	req.Header.Add("Authorization", "Basic "+auth)

	// Execute token request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error issuing API request: %v", err)
	}

	// Check the response was okay
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid API status: %d", resp.StatusCode)
	}

	// Extract the body data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error receiving token response: %v", err)
	}

	// Unmarshal to BusinessNZ token format
	var t BusinessNzToken
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, fmt.Errorf("error decoding token: %v", err)
	}

	// Convert expiry time to work with standard oauth2 tokens
	return &oauth2.Token{
		AccessToken:  t.AccessToken,
		TokenType:    t.TokenType,
		RefreshToken: "",
		Expiry:       time.Now().Add(time.Second * time.Duration(t.ExpiresIn)),
	}, nil

}

// Helper to execute queries on the BusinessNZ API and decode the response via JSON
func (a *Base) Query(path string, data interface{}) error {

	// Setup HTTP request
	req, err := http.NewRequest("GET", a.baseUrl+path, nil)
	if err != nil {
		return fmt.Errorf("error creating request object: %v", err)
	}

	if a.debug {
		log.Printf("API request: %+v", req)
	}

	// Issue API request
	resp, err := a.Do(req)
	if err != nil {
		return fmt.Errorf("error looking up entity")
	}

	if a.debug {
		log.Printf("API response: %+v", resp)
	}

	// Check response code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("NZBN api returned error %d: %s", resp.StatusCode, resp.Status)
	}

	// Fetch response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading NZBN response body: %v", err)
	}

	if a.debug {
		log.Printf("Response body: %v", string(body))
	}

	// Decode JSON
	err = json.Unmarshal(body, data)
	if err != nil {
		return fmt.Errorf("error decoding NZBN response: %v", err)
	}

	return nil
}

// Not -quite- standard API token, must be converted to an oauth2.Token
type BusinessNzToken struct {
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint   `json:"expires_in"`
	AccessToken string `json:"access_token"`
}
