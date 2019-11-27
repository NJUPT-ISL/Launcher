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
	Long: `该命令用于登录校园网，你需要把你的账号和使用 base64 加密后的密码
以yaml文件的形式保存在文件中，执行 launcher login 命令将默认读取您 $HOME/.l.yaml 文件并登录南邮校园网。
	如果您没有配置账号 yaml 文件，您可以使用如下命令生成 yaml 文件并手工修改用户名密码：
		launcher gen
	如果您登录校园网想使用自定义的文件路径，您可以使用 -c 参数引导账号文件，例如：
		launcher login -c /root/.l.yaml
	您可以用如下网址将您的密码加密：
		https://base64.supfree.net/
`,
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
		if ChinaNetWifi {
			if op.LoginChinaNetWifi(user, string(pass),op.GetIP()){
				fmt.Println("南邮校园网登录成功！")
			}else {
				fmt.Println("南邮校园网登录失败！请检查您的配置！")
			}
		}else {
			if op.DefaultLogin(user, string(pass)){
				fmt.Println("南邮校园网登录成功！您可以使用 launcher logout 命令登出校园网！")
			}else {
				fmt.Println("南邮校园网登录失败！请检查您的配置！")
			}
		}
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
	loginCmd.PersistentFlags().BoolVar(&ChinaNetWifi, "CNWifi", false,"登录NJUPT-ChinaNet Wifi")

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
		fmt.Println("正在使用配置文件:", viper.ConfigFileUsed())
	}
}