package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

/*
RandAllString  生成随机字符串([a~zA~Z0~9])
lenNum 长度
*/
func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(chars)
	for i := 0; i < lenNum; i++ {
		l := chars[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*
RandIntStr  生成随机数字字符串([0~9])
lenNum 长度
*/
func RandIntStr(lenNum int) string {
	str := strings.Builder{}
	length := 10
	for i := 0; i < lenNum; i++ {
		str.WriteString(chars[52+rand.Intn(length)])
	}
	return str.String()
}

/*
RandString  生成随机字符串(a~zA~Z])
lenNum 长度
*/
func RandString(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteString(chars[rand.Intn(length)])
	}
	return str.String()
}

func UUIDV4() string {
	v4, _ := uuid.NewV4()
	return v4.String()
}

func UUIDV4WithoutLine() string {
	uid, _ := uuid.NewV4()
	return strings.ReplaceAll(uid.String(), "-", "")
}

func UUIDV4WithTimeStamp() string {
	return fmt.Sprintf("%s-%d", UUIDV4WithoutLine(), time.Now().Unix())
}
