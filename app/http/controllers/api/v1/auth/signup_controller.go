// 处理用户身份认证相关逻辑

package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/gohub/app/data/user"
	v1 "github.com/sjxiang/gohub/app/http/controllers/api/v1"
)

// 注册控制器
type SignupController struct {
	v1.BaseAPIController
}


// 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 请求对象
	var phoneExistRequest struct {
		Phone string `json:"phone"`
	} 

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&phoneExistRequest); err != nil {
		
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		// 打印错误信息
		fmt.Println(err.Error())

		// 出错了，中断请求
		return
	} 


	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(phoneExistRequest.Phone),
	})
}