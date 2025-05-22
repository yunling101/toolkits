package text

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GenerateRandomToString(length int) string {
	// 使用当前时间作为随机数种子，创建独立的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var passwd = make([]byte, length, length)
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < length; i++ {
		index := r.Intn(len(str))
		passwd[i] = str[index]
	}
	return string(passwd)
}

func RandomSecret(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letterRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")
	bytes := make([]rune, length)

	for i := range bytes {
		bytes[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(bytes)
}

func GenerateRandomToInt(n int) int {
	var (
		usedNumbers = sync.Map{}
		r           = rand.New(rand.NewSource(time.Now().UnixNano()))
	)
	if n <= 0 {
		return 0
	}

	// 计算随机数的范围
	minValue := 1
	for i := 1; i < n; i++ {
		minValue *= 10
	}
	maxValue := minValue*10 - 1

	// 生成唯一的随机数
	for i := 0; i < 100; i++ { // 最多尝试 100 次
		num := r.Intn(maxValue-minValue+1) + minValue
		if _, exists := usedNumbers.Load(num); !exists {
			usedNumbers.Store(num, true)
			return num
		}
	}

	return 0
}

func GenerateRandomRgbColor() string {
	v := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := int(uint8(v.Intn(255)))
	g := int(uint8(v.Intn(255)))
	b := int(uint8(v.Intn(255)))
	a := int(uint8(v.Intn(255)))
	return fmt.Sprintf("rgba(%v,%v,%v,%v)", r, g, b, a)
}
