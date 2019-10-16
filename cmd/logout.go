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
	op "github.com/NJUPT-ISL/Launcher/pkg/operations"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "登出校园网",
	Long: `该命令用于登出校园网，使用 launcher logout 将退出南邮校园网`,
	Run: func(cmd *cobra.Command, args []string) {
		op.GetOperation()
		fmt.Println("Launcher: 已登出南邮校园网。")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
