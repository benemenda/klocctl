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
	"os"
	"text/template"

	"github.com/benemenda/klocctl/config"

	"github.com/spf13/cobra"
)

const healthTplt = `
klocwork: {{if .}}✔{{else}}✘{{end}}
`

const healthTpltProm = `
prometheus: {{if .}}✔{{else}}✘{{end}}
`

var Prometheus bool

type health struct {
	Klocwork interface{} `json:"klocwork"`
}

type prometheusHealth struct {
	Prometheus interface{} `json:"prometheus"`
}

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Checks the availability of the configured klocwork web server",
	Long:  `Checks the availability of the configured klocwork web server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("health called")
		ok := config.IsHealthy()
		err := template.Must(template.New("health").Parse(healthTplt)).Execute(os.Stdout, ok)
		if err != nil {
			fmt.Println(err)
		}
		if Prometheus {
			ok := config.IsHealthyProm()
			err := template.Must(template.New("prometheusHealth").Parse(healthTpltProm)).Execute(os.Stdout, ok)
			if err != nil {
				fmt.Println(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
	healthCmd.PersistentFlags().BoolVarP(&Prometheus, "prometheus", "p", false, "Also check health of a Prometheus server")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// healthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// healthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
