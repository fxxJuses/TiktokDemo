package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func InputPassToFromPass(inputPass string) string {
	salt := "1a2b3c4d"
	str := strconv.Itoa(int(salt[0])) + strconv.Itoa(int(salt[2])) + inputPass + strconv.Itoa(int(salt[5])) + strconv.Itoa(int(salt[4]))

	return MD5(str)
}

func FromPassToDBPass(fromPass string, salt string) string {
	str := strconv.Itoa(int(salt[0])) + strconv.Itoa(int(salt[2])) + fromPass + strconv.Itoa(int(salt[5])) + strconv.Itoa(int(salt[4]))
	return MD5(str)
}

func InputPassDBPass(inputPass string, salt string) string {
	fromPass := InputPassToFromPass(inputPass)
	dbPass := FromPassToDBPass(fromPass, salt)
	return dbPass
}
