package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ListingOptions struct {
	After  string `json:"after,omitempty"`
	Before string `json:"before,omitempty"`
	Count  int    `json:"count"`
	Limit  int    `json:"limit""`
	Show   string `json:"show""`
}

type Credentials struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	UserAgent    string `json:"user_agent"`
}

const (
	TestConfigFile = "~/.gtr_credentials"
)

func main() {
	//fmt.Println(frontPage())
	fmt.Println(redditAuth())
}

func frontPage() (string, error) {
	url := "http://www.reddit.com/best.json?"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for %s: %v", url, err)
	}

	// Dodge 429's with this header due to OAuth2
	req.Header.Set("User-agent", "GTR")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	return string(body), nil
}

func redditAuth() (string, error) {

	url := "https://www.reddit.com/api/v1/authorize.compact?"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for %s: %v", url, err)
	}

	//Create the required query string
	q := req.URL.Query()
	q.Add("state", "testing")
	q.Add("response_type", "200")
	q.Add("scope", "identity")
	q.Add("client_id", "GTR")
	q.Add("redirect_url", "https://www.google.com")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	// Dodge 429's with this header due to OAuth2
	req.Header.Set("User-agent", "GTR")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body), nil


	// NEXT STEP *********

	url2 := "https://ssl.reddit.com/api/v1/access_token"
	req2, err := http.NewRequest("POST", url2, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for %s: %v", url, err)
	}

	//Create the required query string
	q2 := req2.URL.Query()
	q2.Add("state", "testing")
	q2.Add("grant_type", "200")
	q2.Add("scope", "identity")
	q2.Add("client_id", "GTR")
	q2.Add("redirect_url", "https://www.google.com")
	q2.Add("code", "wack")
	req2.URL.RawQuery = q2.Encode()

	fmt.Println(req2.URL.String())

	// Dodge 429's with this header due to OAuth2
	req2.Header.Set("User-agent", "GTR")
	//BasicAuth Needs to be in this request
	res2, err := http.DefaultClient.Do(req2)
	if err != nil {
		return "", err
	}
	defer res2.Body.Close()

	body2, _ := ioutil.ReadAll(res2.Body)

	return string(body2), nil


	// Bearer Token Auth
	url3 := "https://oauth.reddit.com/api/v1/me"
	req3, err := http.NewRequest("GET", url3, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for %s: %v", url, err)
	}

	// Dodge 429's with this header due to OAuth2
	req3.Header.Set("User-agent", "GTR")
	req3.Header.Set("Authroization", "BearerToken")
	//BasicAuth Needs to be in this request
	res3, err := http.DefaultClient.Do(req3)
	if err != nil {
		return "", err
	}
	defer res3.Body.Close()

	body3, _ := ioutil.ReadAll(res3.Body)

	return string(body3), nil
}

func tokenRefresh()(string, error){
	url := "https://ssl.reddit.com/api/v1/access_token"
	tokenReq, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for %s: %v", url, err)
	}

	//Create the required query string
	q2 := tokenReq.URL.Query()
	q2.Add("state", "testing")
	q2.Add("grant_type", "200")
	q2.Add("scope", "identity")
	q2.Add("client_id", "GTR")
	q2.Add("redirect_url", "https://www.google.com")
	q2.Add("code", "wack")
	tokenReq.URL.RawQuery = q2.Encode()

	fmt.Println(tokenReq.URL.String())

	// Dodge 429's with this header due to OAuth2
	tokenReq.Header.Set("User-agent", "GTR")
	//BasicAuth Needs to be in this request
	res2, err := http.DefaultClient.Do(tokenReq)
	if err != nil {
		return "", err
	}
	defer res2.Body.Close()

	body2, _ := ioutil.ReadAll(res2.Body)

	return string(body2), nil
}