package text

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Md5Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateMd5ToString(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
