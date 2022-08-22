package user

import (
	"github.com/sjxiang/gohub/pkg/database"
	"github.com/sjxiang/gohub/pkg/hash"
)

// IsEmailExist 判断 Email 是否已经被注册
func IsEmailExist(email string) bool {
	var count int64

	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}


// IsPhoneExist 判断手机号是否已经被注册
func IsPhoneExist(phone string) bool {
	var count int64

	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}


// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
} 


// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}