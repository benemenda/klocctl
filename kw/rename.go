package kw

import "fmt"

/*
Implements:
update_build
Update a build.
Example: curl --data "action=update_build&user=myself&name=build_1&new_name=build_03_11_2011" http://127.0.0.1:8090/review/api
project*
project name
name*
build name
new_name
new build name
keepit
whether this build will be deleted by the auto-delete build feature (true|false)
*/
func renameBuild(args []string) {
	project := args[0]
	oldName := args[1]
	newName := args[2]
	data, klocworkURL := formBaseRequest("builds")

	data.Set("action", "update_build")
	data.Set("project", project)
	data.Set("name", oldName)
	data.Set("new_name", newName)
	_, body := sendRequest(klocworkURL, data)

	if body != nil {
		fmt.Println("Done.")
	}
}

func renameBuilds(projectNames []string) {
	for _, projectName := range projectNames {
		data, klocworkURL := formBaseRequest("builds")
		data.Set("action", "builds")
		data.Set("project", projectName)

		fmt.Println("Retrieving builds for project " + projectName)

		//Send it
		_, body := sendRequest(klocworkURL, data)

		//Get the list of builds
		buildNames := getProjects(body, "builds")
		if buildNames != nil {
			for _, buildName := range buildNames {
				data.Set("action", "update_build")
				data.Set("name", buildName)
				data.Set("new_name", (buildName + ".old"))

				fmt.Println("Project: " + projectName)
				fmt.Println("Renaming build " + buildName + " to new name: " + (buildName + ".old"))

				_, body := sendRequest(klocworkURL, data)
				if body != nil {
				}
			}
		}
	}

}
