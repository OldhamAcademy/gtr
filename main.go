package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://www.reddit.com/best"
	req, _ := http.NewRequest("GET", url, nil)

	// Dodge 429's with this header due to OAuth2
	req.Header.Set("User-agent", "GTR")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}