package Hissec

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetMd5Key() string {
	time.Sleep(time.Microsecond)
	rand.Seed(time.Now().UnixNano())
	key := md5.Sum([]byte(strconv.Itoa(rand.Int())))
	return fmt.Sprintf("%x", key)
}
