package config

import "net/http"
import "fmt"

//IsHealthy return Health clair result
func IsHealthy() bool {
	fmt.Println("requesting health on: " + url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("requesting klocwork health: %v", err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false
	}

	return true
}
