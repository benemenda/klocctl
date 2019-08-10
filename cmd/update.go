/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"klocctl/kw"
	"os"

	"github.com/spf13/cobra"
)

var project string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update some klocwork server resource.",
	Long: `update_build
	update_defect_type
	update_group
	update_module
	update_project
	update_role_assignment
	update_role_permissions
	update_status
	update_view`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		if len(args) < 1 {
			fmt.Printf("klocctl: \"update\" requires a minimum of 1 resource'\n")
			os.Exit(1)
		}
		if args[0] == "status" {
			if len(args[1:]) < 2 {
				fmt.Printf("klocctl: \"update status\" requires a minimum of [project] [id] [status]'\n")
				os.Exit(1)
			}
			kw.ReceiveRequest("update", "status", args[1:])
		}
	},
}

var cmdStatus = &cobra.Command{
	Use:   "status [project] [id] [status]",
	Short: "update the status of an issue on the klocwork server",
	Long:  `update the status of an issue on the klocwork server`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("update status called")
		kw.ReceiveRequest("update", "status", args)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.
	cmdStatus.Flags().StringP("project", "p", project, "klocwork project")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")
	//updateCmd.PersistentFlags().String("status", "s", "update status of an issue")
	updateCmd.AddCommand(cmdStatus)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
