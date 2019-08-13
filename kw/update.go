package kw

import "fmt"

/*
Change the status, owner, and comment, or alternatively set the bug tracker id of issues.
Example: curl --data "action=update_status&user=myself&project=my_project&ids=ids_list&status=new_status&comment=new_comment&owner=new_owner" http://127.0.0.1:8090/review/api
project*
project name
ids*
comma seperated list of ids to change
status
new status to set
comment
new comment to set
owner
new owner to set
bug_tracker_id
new bug tracker id to set

./klocwork update status [project] [ids] [status] [comment] [owner]
*/
func updateStatus(args []string) {
	data, klocworkURL := formBaseRequest("update_status")
	data.Set("project", args[0])
	data.Set("ids", args[1])
	data.Set("status", args[2])

	sendRequest(klocworkURL, data)

}

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
func updateBuild(args []string) {
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
