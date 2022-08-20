// 存放程序的所有子命令

package cmd

import (
	"os"

	"github.com/sjxiang/gohub/pkg/util"
	"github.com/spf13/cobra"
)

// Env 存储全局选项 --env 的值
var Env string


// RegisterFlags 注册全局选项（flag）
func RegisterFlags(rootCmd *cobra.Command) {

	// p 指向 flags 值
	// name 参数名
	// shorthand 缩写
	// value 缺省值
	// usage 详情 tips
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "加载 .env 文件，例：--env=testing，则加载 .env.testing 文件" )
	
}


// RegisterDefaultCmd 注册默认命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {

	cmd, _ , err := rootCmd.Find(os.Args[1:])
	firstArg := util.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
	
}

/*
补充

$ go run main.go --h
os.Args[:] 即 [/tmp/go-build1757847827/b001/exe/main --h]

*/