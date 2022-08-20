// 处理请求数据和表单验证、

package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/sjxiang/gohub/pkg/response"
	"github.com/thedevsaddam/govalidator"
)

// 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string


func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {

	// 1. 解析请求，支持 JSON 数据
	if err := c.ShouldBindJSON(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		// 	"error": err.Error,
		// })
		
		// fmt.Println(err.Error())
		return false
	}

	// 2. 表单验证
	errs := handler(obj, c)  // 回调函数，钩子 包装 requests 包下面的 ValidateSignupEmailExist 等方法

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message": "请求验证不通过，具体请查看 errors",
		// 	"errors": errs,
		// })
		return false
	}

	return true
}


// === 

type SignupPhoneExistRequest struct {
	Phone string  `json:"phone,omitempty" valid:"phone"`
}

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	
	// 自定义规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称为 phone",  // Ps. 此处 “:”，为英文符号
			"digits:手机号长度必须为 11 位的数字",
		},
	} 

	return validate(data, rules, messages)
}



// ===

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}


func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	// 开始验证
	return validate(data, rules, messages)
}


// 重构
func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	// 配置选项
	opts := govalidator.Options {
		Data: data,
		Rules: rules,
		TagIdentifier: "valid",  // 模型中的 Struct 标签标识符
		Messages: messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}