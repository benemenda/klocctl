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

package kw

import (
	"fmt"
	"io/ioutil"
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

type Vertex struct {
	Lat, Long float64
}

func checkProjectExists(project string, projectNames []string) bool {
	for _, proj := range projectNames {
		if project == proj {
			return true
		}
	}
	return false
}

/*
Receives some CLI-based request requiring data from the Klocwork server.
*/
func ReceiveRequest(verb, command string, args []string) []string {
	var returnValue []string
	data, klocworkURL := formBaseRequest("projects")
	_, body := sendRequest(klocworkURL, data)

	switch verb {
	case "get":
		switch command {
		case "projects":
			getProjects(body, "projects")
		case "builds":
			getBuilds(args)
		case "issues":
			getIssues(args)
		}
	case "update":
		projectNames := getProjects(body, "projects")
		switch command {
		case "status":
			for _, proj := range projectNames {
				//valid project so try to update status
				if args[0] == proj {
					updateStatus(args)
				}
			}
		case "build":
			//TODO: checkBuildExists
			if checkProjectExists(args[0], projectNames) {
				updateBuild(args)
			}
		}
	}
	return returnValue
}

func formBaseRequest(command string) (url.Values, string) {
	//Form the URL
	protocol, host, port, user, ltoken := viper.GetString("klocctl.protocol"),
		viper.GetString("klocctl.host"),
		viper.GetString("klocctl.port"),
		viper.GetString("klocctl.user"),
		viper.GetString("klocctl.ltoken")

	fmt.Printf("%v", host)
	var klocworkURL = protocol + "://" + host + ":" + port + "/review/api"

	//Create the request
	data := url.Values{}
	data.Set("action", command)
	// data.Set("user", "<USERNAME>")
	// data.Set("ltoken", "<LTOKEN>")
	data.Set("user", user)
	data.Set("ltoken", ltoken)

	return data, klocworkURL
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
