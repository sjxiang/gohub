// 处理用户身份认证相关逻辑

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/gohub/app/data/user"
	v1 "github.com/sjxiang/gohub/app/http/controllers/api/v1"
	"github.com/sjxiang/gohub/app/requests"
)

// 注册控制器
type SignupController struct {
	v1.BaseAPIController
}


// 检测手机号是否已注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 测试 
	// panic("这是 panic 测试")

	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}


// 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})

}



// === 
// 重构

// // 检测手机号是否已注册
// func (sc *SignupController) IsPhoneExist(c *gin.Context) {

// 	// 初始化请求对象 
// 	request := requests.SignupPhoneExistRequest{}

// 	// 解析 JSON 请求
// 	if err := c.ShouldBindJSON(&request); err != nil {
		
// 		// 解析失败，返回 422 状态码和错误信息
// 		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
// 			"error": err.Error(),
// 		})

// 		// 打印错误信息
// 		fmt.Println(err.Error())

// 		// 出错了，中断请求
// 		return
// 	} 


// 	// 表单验证
// 	errs := requests.ValidateSignupPhoneExist(&request, c) 

// 	// errs 返回长度，等于 0，即通过；大于 0，即有错误发生
// 	if len(errs) > 0 {
		
// 		// 验证失败，返回 422 状态码和错误信息
// 		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
// 			"errors": errs,
// 		})
// 		return
// 	}

// 	// 检查数据库并返回响应
// 	c.JSON(http.StatusOK, gin.H{
// 		"exist": user.IsPhoneExist(request.Phone),
// 	})
// }


// // 检测邮箱是否已注册
// func (sc *SignupController) IsEmailExist(c *gin.Context) {

	
// 	// 初始化请求对象
// 	request := requests.SignupEmailExistRequest{}

// 	// 解析 JSON 对象
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
// 			"error": err.Error(),
// 		})
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// 表单验证
// 	errs := requests.ValidateSignupEmailExist(&request, c)
// 	if len(errs) > 0 {
// 		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
// 			"errors": errs,
// 		})
// 		return
// 	}

// 	// 检查数据库并返回
// 	c.JSON(http.StatusOK, gin.H{
// 		"exist": user.IsEmailExist(request.Email),
// 	})

// }

