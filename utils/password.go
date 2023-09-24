package utils

import "golang.org/x/crypto/bcrypt"

// 密码加密处理
func EncryptPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(fromPassword), err
}

// 密码验证
func ParsingPassword(encryptPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
