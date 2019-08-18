// Example usage of the client. Build with:
// go build -o httpbin-client
// and execute the binary to run HTTP requests towards httpbin.org
package main

import (
	"httpbin_client/client"
)

func main() {
	hostname := "https://httpbin.org"
	apiClient := client.NewClient(hostname)
	apiClient.Get()
	apiClient.Post()
	apiClient.Delete()
	apiClient.Put()
}
