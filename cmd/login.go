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
	"encoding/base64"
	"fmt"
	op "github.com/NJUPT-ISL/Launcher/pkg/operations"
	ya "github.com/NJUPT-ISL/Launcher/pkg/yaml"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登录校园网",
	Long: `用于登录校园网，你需要把你的账号和密码告知 Launcher, Launcher 会帮助你登录你的校园网。`,
	Run: func(cmd *cobra.Command, args []string) {
		user, B64Pass, err := ya.ReadYaml(cfgFile)
		if err != nil {
			fmt.Printf("读取错误：%v",err)
			return
		}
		pass, err := base64.StdEncoding.DecodeString(B64Pass)
		if err != nil {
			fmt.Printf("读取错误：%v",err)
			return
		}
		op.PostOperation(user, string(pass))
	},

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	loginCmd.PersistentFlags().StringVarP(&cfgFile, "config","c", os.Getenv("HOME")+"/.l.yaml", "配置文件路径")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".Launcher" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".launcher")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}