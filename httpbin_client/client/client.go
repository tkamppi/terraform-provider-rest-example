package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname string
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string) *Client {
	return &Client{
		hostname: hostname,
	}
}

// Get makes a HTTP GET request to the hostname/get URL
func (s *Client) Get() error {

	url := s.hostname + "/get"
	request, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("HTTP Response Status:", response.StatusCode)
		fmt.Println("HTTP Body:")
		fmt.Println(string(data))
	}
	return nil
}

// Delete makes a HTTP DELETE request to the hostname/delete URL
func (s *Client) Delete() error {

	url := s.hostname + "/delete"
	request, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("HTTP Response Status:", response.StatusCode)
		fmt.Println("HTTP Body:")
		fmt.Println(string(data))
	}
	return nil
}

// Post makes a HTTP POST request to the hostname/post URL
func (s *Client) Post() error {

	url := s.hostname + "/post"
	request, _ := http.NewRequest("POST", url, nil)
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("HTTP Response Status:", response.StatusCode)
		fmt.Println("HTTP Body:")
		fmt.Println(string(data))
	}
	return nil
}

// Put makes a HTTP PUT request to the hostname/put URL
func (s *Client) Put() error {

	url := s.hostname + "/put"
	request, _ := http.NewRequest("PUT", url, nil)
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("HTTP Response Status:", response.StatusCode)
		fmt.Println("HTTP Body:")
		fmt.Println(string(data))
	}
	return nil
}
