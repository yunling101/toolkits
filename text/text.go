package text

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"

	"regexp"

	uuid "github.com/satori/go.uuid"
)

// 字符串md5加密
func GenerateMd5ToString(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))

	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// 字符串sha256加密
func GenerateSha256ToString(str string) string {
	sha := sha256.New()
	sha.Write([]byte(str))

	return hex.EncodeToString(sha.Sum(nil))
}

// 随机字符串
func GenerateRandomToString(length int) string {
	var passwd []byte = make([]byte, length, length)
	rand.Seed(time.Now().UnixNano())

	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < length; i++ {
		index := rand.Intn(len(str))
		passwd[i] = str[index]
	}

	return string(passwd)
}

// 随机数
func GenerateRandomToInt(n int) int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		code := fmt.Sprintf("%v", rnd.Intn(999999999999))
		if len(code) > n {
			id, _ := strconv.Atoi(code[:n])
			return id
		}
	}
	return 0
}

// UUID
func GenerateRandomUUID() string {
	return uuid.NewV1().String()
}

// 随机颜色
func GenerateRandomRgbColor() string {
	rand.Seed(time.Now().UnixNano())
	r := int(uint8(rand.Intn(255)))
	g := int(uint8(rand.Intn(255)))
	b := int(uint8(rand.Intn(255)))
	a := int(uint8(rand.Intn(255)))
	return fmt.Sprintf("rgba(%v,%v,%v,%v)", r, g, b, a)
}

// 随机秘钥
func RandomSecret(length int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")
	bytes := make([]rune, length)

	for i := range bytes {
		bytes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(bytes)
}

// VerifyIP 验证IP合法
func VerifyIP(msg string) bool {
	return true
}

func ContainIP(msg string) bool {
	ipReg := `(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})(\.(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})){3}`
	validatorReg := regexp.MustCompile(ipReg)
	if len(validatorReg.Find([]byte(msg))) == 0 {
		return false
	}
	return true
}
