package cryptoutils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"

	"github.com/sony/sonyflake/v2"
	"golang.org/x/crypto/bcrypt"
)

func Sha512(value string) string {
	sum := sha512.Sum512([]byte(value))
	return hex.EncodeToString(sum[:])
}

func getMachineID() (int, error) {
	return 1, nil
}

// RandomSonyflake 雪花ID
func RandomSonyflake() int64 {
	settings := sonyflake.Settings{
		MachineID: getMachineID, // 必须提供这个函数
	}

	sf, _ := sonyflake.New(settings)
	id, _ := sf.NextID()
	return id
}

// GenerateToken 生成 Token 用于登陆
func GenerateToken() string {
	b := make([]byte, 32) // 256bit
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

// HashPassword 登陆密码加密
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPassword 验证登陆密码
func CheckPassword(hash, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
