package helper

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/debidarmawan/debozero-core/constants"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), constants.HashingCost)
	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckOldPasswordHash(password string, hash string) bool {
	return strings.ToUpper(GetMD5Hash(password)) == hash
}

func GetMD5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
