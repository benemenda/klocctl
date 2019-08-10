package kw

import "fmt"

func renameBuilds(projectNames []string) {
	for _, projectName := range projectNames {
		data, klocworkUrl := formBaseRequest("builds")
		data.Set("action", "builds")
		data.Set("project", projectName)

		fmt.Println("Retrieving builds for project " + projectName)

		//Send it
		_, body := sendRequest(klocworkUrl, data)

		//Get the list of builds
		buildNames := getProjects(body, "builds")
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
