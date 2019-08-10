package kw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

//Internal function to get a list of project names
// Input: Web API JSONResponse []byte
// Output: List of project names []string
func getProjects(aJSONResponse []byte, aType string) []string {
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

func getBuilds(projectNames []string) {

	for _, projectName := range projectNames {
		data, klocworkURL := formBaseRequest("builds")
		//data.Set("action", "builds")
		data.Set("project", projectName)
		fmt.Println("Retrieving builds for project " + projectName)
		//Send it
		_, body := sendRequest(klocworkURL, data)
		buildNames := getProjects(body, "builds")
		fmt.Println("Project: " + projectName)
		fmt.Println("Builds: ")
		for _, buildName := range buildNames {
			fmt.Println(buildName)
		}
	}
	return
}

/*
Implements:
search
Retrieve the list of detected issues.
Example: curl --data "action=search&user=myself&project=my_project&query=file:MyFile.c" http://127.0.0.1:8090/review/api
project*
project name
query
search query, such as narrowing by file (for example, 'file:MyFile.c')
view
view name
limit
search result limit
summary
include summary record to output stream
*/
func getIssues(args []string) {
	data, klocworkURL := formBaseRequest("search")
	project := args[0]
	data.Set("project", project)
	if len(args) > 1 {
		query := args[1]
		data.Set("query", query)
		fmt.Println("Retrieving issues for project " + project + " matching search query " + data.Get("query"))
	} else {
		fmt.Println("Retrieving all issues for project " + project)
	}
	//Send it
	sendRequest(klocworkURL, data)
}
