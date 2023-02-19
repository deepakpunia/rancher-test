// rancher_test.go

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	// Set up a test server with the Rancher API endpoints
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/v3-public/localProviders/local?action=login" {
			// Validate the login request
			if r.Header.Get("Content-Type") != "application/json" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Expected Content-Type of application/json")
				return
			}
			if r.Header.Get("User-Agent") != "Rancher API Client" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Expected User-Agent of Rancher API Client")
				return
			}
			// Respond with a successful login token
			fmt.Fprintln(w, `{"token":"mytoken"}`)
			return
		}
		// Return a 404 for any other requests
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create a new Rancher API client with the test server URL
	client := NewRancherClient(ts.URL)

	// Log in with a test username and password
	token, err := client.Login("myusername", "mypassword")
	if err != nil {
		t.Errorf("Error logging in: %v", err)
	}
	if token != "mytoken" {
		t.Errorf("Expected token of 'mytoken', got '%s'", token)
	}
}

type RancherClient struct {
	URL string
}

func NewRancherClient(url string) *RancherClient {
	return &RancherClient{
		URL: url,
	}
}

func (c *RancherClient) Login(username, password string) (string, error) {
	req, err := http.NewRequest("POST", c.URL+"/v3-public/localProviders/local?action=login", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Rancher API Client")
	q := req.URL.Query()
	q.Add("username", username)
	q.Add("password", password)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Expected status code of 200, got %d", resp.StatusCode)
	}

	var body struct {
		Token string `json:"token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.Token, nil
}

