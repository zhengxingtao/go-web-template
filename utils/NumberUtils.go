package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var (
	//password security key
	PwSk = "2020"
)

func RandNumber() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

//将key + password + salt拼接起来后做一次md5
func PasswordMd5(password string, salt string) string {
	passwordStr := PwSk + password + salt
	hash := md5.New()
	_, _ = hash.Write([]byte(passwordStr))
	return hex.EncodeToString(hash.Sum(nil))
}
