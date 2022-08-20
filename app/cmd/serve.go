package cmd

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/gohub/bootstrap"
	"github.com/sjxiang/gohub/pkg/console"
	"github.com/sjxiang/gohub/pkg/logger"
	"github.com/spf13/cobra"
)

// CmdServe 代表可用的 web 子命令
var CmdServe = &cobra.Command{
	Use: "serve",
	Short: "Start web server",
	Run: runWeb,
	Args: cobra.NoArgs,
}


func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin 的运行模式，支持 debug、release、test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到程序中的 log
	// 故设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)


	// new 一个 Gin Engine 实例（指针对象，不会被逃逸分析或垃圾回收干掉，尽情配置）
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)


	// 运行服务，指定监听端口为 9090
	err := router.Run(":"+os.Getenv("APP_PORT"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error: " + err.Error())
	}

}