// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	// "strings"
	"github.com/spf13/cobra"
	restapi "testRestApi/facades"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the REST API",
	Long: `Start and serve the REST API`,
	Run: func(cmd *cobra.Command, args []string) {		
		prefix := cmd.Flag("prefix").Value.String()
		host := cmd.Flag("host").Value.String()
		fmt.Printf("serve called %v %v \n", prefix, host)

		restapi.Init(prefix, host)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP("prefix", "", "", "API prefix ex : api")
	serveCmd.Flags().StringP("host", "", "", "host:port ex : :3000")
}
