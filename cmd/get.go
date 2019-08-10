/*
Copyright © 2019 Ben Marsden ben.marsden@emenda.se

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

	"github.com/spf13/cobra"
)

var queryString string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "pull some resource (e.g. projects, builds) from the configured Klocwork server.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var cmdProjects = &cobra.Command{
	Use:   "projects",
	Short: "get the list of available projects",
	Long:  `get the list of available projects`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("get projects called")
		kw.ReceiveRequest("get", "projects", args)
	},
}
var cmdBuilds = &cobra.Command{
	Use:   "builds [project1, project2, ..., projectN]",
	Short: "get the list of builds for specified projects",
	Long:  `get the list of builds for specified projects`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projects := args
		fmt.Printf("get builds called")
		kw.ReceiveRequest("get", "builds", projects)
	},
}
var cmdIssues = &cobra.Command{
	Use:   "issues [project] (query)",
	Short: "get the list of issues for a specified klocwork project",
	Long:  `get the list of issues for a specified klocwork project`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("get issues called")
		kw.ReceiveRequest("get", "issues", args)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(cmdProjects, cmdBuilds, cmdIssues)
	//cmdIssues.PersistentFlags().StringVarP(&queryString, "query", "q", "", "search query to filter issues by")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
