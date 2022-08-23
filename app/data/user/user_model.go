package user


import (
	"github.com/sjxiang/gohub/app/data"
	"github.com/sjxiang/gohub/pkg/database"
	"github.com/sjxiang/gohub/pkg/hash"
)

// User 用户模型
type User struct {
	data.BaseModel

	Name  	 string `json:"name,omitemity"`
	Email	 string `json:"-"`  // 敏感信息，json 转义忽略
	Phone 	 string `json:"-"`
	Password string `json:"-"`

	data.CommonTimestampsField
}



// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
} 


// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}


