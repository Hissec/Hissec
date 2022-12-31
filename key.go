package Hissec

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//GetMd5Key  生成随机的MD5
func GetMd5Key() string {
	time.Sleep(time.Microsecond)
	rand.Seed(time.Now().UnixNano())
	key := md5.Sum([]byte(strconv.Itoa(rand.Int())))
	return fmt.Sprintf("%x", key)
}

// GetVerifyCode 密码算法
func GetVerifyCode(value string) string {
	key := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", key)
}

