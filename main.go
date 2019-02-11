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

	// Dodge 429's with this header due to OAuth2
	req.Header.Set("User-agent", "GTR")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body), nil
}
