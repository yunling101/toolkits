package text

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSha256ToString(str string) string {
	sha := sha256.New()
	sha.Write([]byte(str))
	return hex.EncodeToString(sha.Sum(nil))
}
