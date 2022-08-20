package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sjxiang/gohub/app/cmd"
	"github.com/sjxiang/gohub/bootstrap"
	"github.com/sjxiang/gohub/conf"
	"github.com/sjxiang/gohub/pkg/console"
)



func init() {
	
	// 从配置文件读取配置
	conf.Init()
}


func main() {

	// 应用的主入口，默认调用 cmd.Serve 命令
	var rootCmd = &cobra.Command{
		Use: os.Getenv("APP_NAME"),
		Short: "一个简单的论坛项目，体悟下 crud",
		Long: `默认允许 serve 命令，也可以使用 -h 命令，看看其它子命令`,
		
		// rootCmd 的所有子命令都会执行以下代码，前置
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			
			// 初始化 Logger
			bootstrap.SetupLogger()
		
			// 初始化 DB 
			bootstrap.SetupDB()

			// 初始化缓存
			bootstrap.SetupRedis()

		},
	}


	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
	)

	// // 配置默认运行 web 服务
	// cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	
	// 注册全局参数，--env
	cmd.RegisterFlags(rootCmd) 

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}

