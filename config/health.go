package config

import "net/http"
import "fmt"

//IsHealthy return Health Klocwork result
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

//IsHealthyProm return Health Prometheus result
func IsHealthyProm() bool {
	fmt.Println("requesting health on: " + urlProm)
	response, err := http.Get(urlProm)
	if err != nil {
		fmt.Printf("requesting prometheus health: %v", err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false
	}

	return true
}
