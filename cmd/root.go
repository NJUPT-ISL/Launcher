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
  "github.com/spf13/cobra"
  "os"
)


var (
  cfgFile string
  ChinaNetWifi = false
  )


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "Launcher",
  Short: "Launcher 是 Go 语言开发的南京邮电大学校园网登录器",
  Long: `
██╗      █████╗ ██╗   ██╗███╗   ██╗ ██████╗██╗  ██╗███████╗██████╗ 
██║     ██╔══██╗██║   ██║████╗  ██║██╔════╝██║  ██║██╔════╝██╔══██╗
██║     ███████║██║   ██║██╔██╗ ██║██║     ███████║█████╗  ██████╔╝
██║     ██╔══██║██║   ██║██║╚██╗██║██║     ██╔══██║██╔══╝  ██╔══██╗
███████╗██║  ██║╚██████╔╝██║ ╚████║╚██████╗██║  ██║███████╗██║  ██║
╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
                                                                      
Launcher 是南邮 ISL 实验室基于 Go 语言开发的校园网登录器，
用于服务器接入校园网使用，现已开源：https://github.com/NJUPT-ISL/Launcher
你需要把你的账号和使用 base64 加密后的密码以yaml文件的形式保存在文件中，默认读
取您 $HOME 下的 .l.yaml 文件。
	如果您第一次使用，您可以使用如下命令根据提示生成配置文件:
		launcher gen -h
	您可以使用如下命令根据提示登录校园网:
		launcher login -h
    您可以使用如下命令根据提示登出校园网:
		launcher logout -h
`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}





