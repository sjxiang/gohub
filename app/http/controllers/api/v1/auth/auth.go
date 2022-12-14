// 授权相关逻辑

package auth

import (
	"errors"

	"github.com/sjxiang/gohub/app/data/user"
)



func Attempt(email string, password string) (user.User, error) {

	userModel := user.GetByMulti(email) 
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil 
}

