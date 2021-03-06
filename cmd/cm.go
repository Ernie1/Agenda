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

	"github.com/spf13/cobra"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetString("participator")
		st, _ := cmd.Flags().GetString("starttime")
		et, _ := cmd.Flags().GetString("endtime")
		fmt.Printf("cm called with args: title=%s participators=%s st=%s et=%s\n", title, participators, st, et)

	},
}

func init() {
	rootCmd.AddCommand(cmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmCmd.Flags().StringP("title","t","","title of meeting")
	cmCmd.MarkFlagRequired("title")
	
	cmCmd.Flags().StringP("participator","p","","participators of eeting")
	cmCmd.MarkFlagRequired("participator")
	
	cmCmd.Flags().StringP("starttime","s","","start time")
	cmCmd.MarkFlagRequired("starttime")

	cmCmd.Flags().StringP("endtime","e","","end time")
	cmCmd.MarkFlagRequired("endtime")
}
