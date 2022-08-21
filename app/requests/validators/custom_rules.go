// 存放自定义规则和验证器

package validators

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sjxiang/gohub/pkg/database"
	"github.com/thedevsaddam/govalidator"
)

// 注册自定义表单验证规则
func init() {

	// 自定义规则 not_exists，验证请求数据必须不存在与数据库中
	// 常用于保证数据库某个字段的值唯一，e.g. 用户名、邮箱、手机号、或者分类的名称 
	// not_exists 参数可以有两种，一种是 2 个参数，另一种是 3 个参数
	// not_exists:user,email 检查数据库表里是否存在同一条信息
	// not_exists:user,email,32 排除用户 id 为 32 的用户 

	govalidator.AddCustomRule("not_exists", func(field, rule, message string, value interface{}) error {

		// 移除前缀
		// 切割
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第 2 个参数，表名称，e.g. users
		tableName := rng[0]

		// 第 2 个参数，字段名称，e.g. email 或者 phone
		dbField := rng[1]

		// 第 3 个参数，排除 ID
		var expectID string
		if len(rng) > 2 {
			expectID = rng[2]
		}

		// 用户请求过来的数据
		requestValue := value.(string)

		// 拼接 SQL
		query := database.DB.Table(tableName).Where(dbField+" = ?", requestValue)

		// 如果传参第 3 个参数，加上 SQL Where 过滤
		if len(expectID) > 0 {
			query.Where("id != ?", expectID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		// 验证不通过，数据库能找到对应的数据
		if count != 0 {

			// 如果有自定义错误消息
			if message != "" {
				return errors.New(message)
			}

			// 默认的错误消息
			return fmt.Errorf("%v 已被占用", requestValue)
		} 

		// 验证通过
		return nil	
	})
}