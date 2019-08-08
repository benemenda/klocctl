/*
MIT License

Copyright (c) [year] [fullname]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Emenda 2019
Author: Andreas Lärfors
*/

/// N.B.!!!!
/// This script is very much work-in-progress!
/// You will need to modify it to do what you want
/// TODO: Command-line parsing and everything parameterised

package kwservertool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Result interface {
	GetName() string
}

//Project result template
type Project struct {
	Id, Name string
}

func (pr *Project) GetName() string {
	return pr.Name
}

//Build result template
type Build struct {
	Id   int
	Name string
}

func (bl *Build) GetName() string {
	return bl.Name
}

////
type Vertex struct {
	Lat, Long float64
}

////
func getBuilds(projectNames []string) {

	for _, projectName := range projectNames {
		data, klocworkUrl := formRequest("builds")
		//data.Set("action", "builds")
		data.Set("project", projectName)
		fmt.Println("Retrieving builds for project " + projectName)
		//Send it
		_, body := sendRequest(klocworkUrl, data)
		buildNames := getNames(body, "builds")
		fmt.Println("Project: " + projectName)
		fmt.Println("Builds: ")
		for _, buildName := range buildNames {
			fmt.Println(buildName)
		}
	}
	return
}

func renameBuilds(projectNames []string) {
	for _, projectName := range projectNames {
		data, klocworkUrl := formRequest("builds")
		data.Set("action", "builds")
		data.Set("project", projectName)

		fmt.Println("Retrieving builds for project " + projectName)

		//Send it
		_, body := sendRequest(klocworkUrl, data)

		//Get the list of builds
		buildNames := getNames(body, "builds")
		if buildNames != nil {
			for _, buildName := range buildNames {
				data.Set("action", "update_build")
				data.Set("name", buildName)
				data.Set("new_name", (buildName + ".old"))

				fmt.Println("Project: " + projectName)
				fmt.Println("Renaming build " + buildName + " to new name: " + (buildName + ".old"))

				_, body := sendRequest(klocworkUrl, data)
				if body != nil {
				}
			}
		}
	}

}

/*
Receives some CLI-based request requiring data from the Klocwork server.
Should answer this request, at the moment only returns project names on the server.
*/
func ReceiveRequest(verb, command string, args []string) []string {
	var returnValue []string
	data, klocworkUrl := formRequest("projects")
	_, body := sendRequest(klocworkUrl, data)

	switch verb {
	case "get":
		switch command {
		case "projects":
			returnValue = getNames(body, "projects")
		case "builds":
			projectNames := getNames(body, "projects")
			getBuilds(projectNames)
		}
	case "rename":
		projectNames := getNames(body, "projects")
		switch command {
		case "builds":
			if projectNames != nil {
				renameBuilds(projectNames)
			}
		//TODO
		case "projects":
			fmt.Println("Renaming projects not yet implemented.")
		}
	}
	return returnValue
}

func formRequest(command string) (url.Values, string) {
	//Form the URL
	protocol, host, port, user, ltoken := viper.GetString("klocctl.protocol"),
		viper.GetString("klocctl.host"),
		viper.GetString("klocctl.port"),
		viper.GetString("klocctl.user"),
		viper.GetString("klocctl.ltoken")

	fmt.Printf("%v", host)
	var klocworkUrl = protocol + "://" + host + ":" + port + "/review/api"

	//Create the request
	data := url.Values{}
	data.Set("action", command)
	// data.Set("user", "<USERNAME>")
	// data.Set("ltoken", "<LTOKEN>")
	data.Set("user", user)
	data.Set("ltoken", ltoken)

	return data, klocworkUrl
}

/*
Request send request for data to the KW server
*/
func Request() {
	//Form the URL
	protocol, host, port, user, ltoken := viper.GetString("klocctl.protocol"),
		viper.GetString("klocctl.host"),
		viper.GetString("klocctl.port"),
		viper.GetString("klocctl.user"),
		viper.GetString("klocctl.ltoken")

	fmt.Printf("%v", host)
	var klocworkUrl = protocol + "://" + host + ":" + port + "/review/api"

	//Create the request
	data := url.Values{}
	data.Set("action", "projects")
	// data.Set("user", "<USERNAME>")
	// data.Set("ltoken", "<LTOKEN>")
	data.Set("user", user)
	data.Set("ltoken", ltoken)

	//Send it
	_, body := sendRequest(klocworkUrl, data)

	//Get the list of projects
	projectNames := getNames(body, "projects")
	if projectNames != nil {
		for _, projectName := range projectNames {

			data.Set("action", "builds")
			data.Set("project", projectName)

			fmt.Println("Retrieving builds for project " + projectName)

			//Send it
			_, body := sendRequest(klocworkUrl, data)

			//Get the list of builds
			buildNames := getNames(body, "builds")
			if buildNames != nil {
				for _, buildName := range buildNames {
					data.Set("action", "update_build")
					data.Set("name", buildName)
					data.Set("new_name", (buildName + ".old"))

					fmt.Println("Project: " + projectName)
					fmt.Println("Renaming build " + buildName + " to new name: " + (buildName + ".old"))

					_, body := sendRequest(klocworkUrl, data)
					if body != nil {
					}
				}
			}
		}
	}

}

//Internal function to get a list of project names
// Input: Web API JSONResponse []byte
// Output: List of project names []string
func getNames(aJSONResponse []byte, aType string) []string {
	var result []string
	dec := json.NewDecoder(bytes.NewReader(aJSONResponse))
	for {

		//Some variables we will need
		var res Result
		var err error

		switch aType {
		case "projects":
			var doc Project
			err = dec.Decode(&doc)
			res = &doc
		case "builds":
			var doc Build
			err = dec.Decode(&doc)
			res = &doc
		default:
			log.Fatal("No implementation for JSON processing of type: " + aType + ". Exiting.")
		}

		//err := dec.Decode(&doc)
		if err == io.EOF {
			//fmt.Printf("EOF")
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if res != nil {
			//fmt.Println(res.GetName())
			result = append(result, res.GetName())
		} else {
			fmt.Println("WARNING: res is nil - error in JSON decoding loop?")
			break
		}
	}
	return result
}

//Internal function to send a request to the Klocwork server
func sendRequest(aUrl string, aData url.Values) (*http.Response, []byte) {
	//Build the request
	req, err := http.NewRequest("POST", aUrl, strings.NewReader(aData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(aData.Encode())))

	//Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	//Print the response
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(body))

	return resp, body
}
