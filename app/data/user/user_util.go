package user

import "github.com/sjxiang/gohub/pkg/database"

// 判断 Email 是否已经被注册
func IsEmailExist(email string) bool {
	var count int64

	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}


// 判断手机号是否已经被注册
func IsPhoneExist(phone string) bool {
	var count int64

	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

