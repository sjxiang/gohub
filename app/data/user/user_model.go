package user


import "github.com/sjxiang/gohub/app/data"


// User 用户模型
type User struct {
	data.BaseModel

	Name  	 string `json:"name,omitemity"`
	Email	 string `json:"-"`  // 敏感信息，json 转义忽略
	Phone 	 string `json:"-"`
	Password string `json:"-"`

	data.CommonTimestampsField
}


