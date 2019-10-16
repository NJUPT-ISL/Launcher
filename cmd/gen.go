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
	"os"

	"github.com/spf13/cobra"
)


// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成配置文件模版",
	Long: `该命令用于生成一个配置文件模版，您只需要将生成的模版文件中的 youruser 改为您的校园网账号，
同时将 b64pass 修改为通过 base64 加密过后的校园网密码即可，执行 launcher gen 命令将默认在 $HOME 位置生成 .l.yaml 文件。
	如果您需要修改生成文件的位置，你可以使用 -f 选项修改创建文件的位置,例如:
		launcher gen -f /root/.l.yaml
	您可以用如下网址将您的密码加密：
		https://base64.supfree.net/`,
	Run: func(cmd *cobra.Command, args []string) {

		f, err := os.Create(cfgFile)
		if err != nil {
			fmt.Println("生成配置文件错误:")
			fmt.Println(err)
			_ = f.Close()
			return
		}
		var exampleYaml = `account:
  user: youruser
  password: b64pass
`
		_, err = f.WriteString(exampleYaml)
		if err != nil {
			fmt.Println(err)
			_ = f.Close()
			return
		}
		fmt.Println("配置文件生成成功!配置文件位于:"+cfgFile)
		fmt.Println(`您只需要将生成的模版文件中的 youruser 改为您的校园网账号，同时将 b64pass 
修改为通过 base64 加密过后的校园网密码即可，此命令将默认在 $HOME 位置生成 .l.yaml 文件。
	如果您需要修改生成文件的位置，你可以使用 -f 选项修改创建文件的位置,例如:
		launcher gen -f /root/.l.yaml
	您可以用如下网址将您的密码加密：
		https://base64.supfree.net/`)
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}



func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.
	loginCmd.PersistentFlags().StringVarP(&cfgFile, "file","f", os.Getenv("HOME")+"/.l.yaml", "生成的配置文件路径")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
