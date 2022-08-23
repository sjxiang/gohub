package user

import (
	"github.com/sjxiang/gohub/pkg/database"
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


// GetByEmail 通过邮箱来获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}


// GetByMulti 通过 『手机号码 / 邮箱 / 用户名』来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.Where("phone = ?", loginID).Or("email = ?", loginID).Or("name = ?", loginID).First(&userModel)
	return
}



