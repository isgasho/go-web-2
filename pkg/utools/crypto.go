package utools

import "golang.org/x/crypto/bcrypt"

// CryptoPassword 密码加密
func CryptoPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// ComparePassword 密码对比
func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
