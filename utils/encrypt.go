package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/snluu/uuid"
	"golang.org/x/crypto/bcrypt"
)

const randNum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const formatForUuid = "1870747d-b26c-4507-9518-1ca62bc66e5d"

func GetMD5(param string) string {
	instance := md5.New()
	instance.Write([]byte(param))
	return hex.EncodeToString(instance.Sum(nil))
}

func GetSHA1(param string) string {
	instance := sha1.New()
	instance.Write([]byte(param))
	return hex.EncodeToString(instance.Sum(nil))
}

func GetPasswordSalt() string {
	return GetRandomString(8)
}

func GetRandomString(length int) string {
	bytes := []byte(randNum)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetUuid() (result string) {
	uuids := uuid.MustFromStr(formatForUuid)
	result = uuids.Hex()
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
