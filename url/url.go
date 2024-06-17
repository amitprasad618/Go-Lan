package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Will Learn abut Uniform Resourse Locator (URL)")

	myURL := "https://amitprasad.com:8080/path/to/resourse/?key=40440&value=71881"

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("error in Parsing URL:", err)
		return
	}

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)

	//modifing URL
	parsedURL.Scheme = "http"
	parsedURL.Host = "newHost.com"
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "key=newValue"

	fmt.Println("----------------------------------------------------------")
	//Constructing a URL string from a URL Object
	newURL := parsedURL.String()
	fmt.Println("Modified URL: ", newURL)
}
