package tools

import "golang.org/x/crypto/bcrypt"

type PwdTool struct {
}

// GeneratePassword 生成密码
func GeneratePassword(rawPassword string) (s string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hash)
}

// CheckPassword 密码校验
func CheckPassword(hashPassword string, rawPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}
