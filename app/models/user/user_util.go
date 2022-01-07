package user

import "huango/pkg/database"

// IsEmailExits 判断 Email 已被注册
func IsEmailExits(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExits 判断 手机号 已被注册
func IsPhoneExits(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
