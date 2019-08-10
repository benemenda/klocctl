package kw

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
	data, klocworkUrl := formBaseRequest("update_status")
	data.Set("project", args[0])
	data.Set("ids", args[1])
	data.Set("status", args[2])

	sendRequest(klocworkUrl, data)

}
