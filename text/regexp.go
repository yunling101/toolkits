package text

import (
	"regexp"
	"strings"
)

func IsMatch(val, pattern string) bool {
	match, err := regexp.Match(pattern, []byte(val))
	if err != nil {
		return false
	}

	return match
}

func IsMail(val string) bool {
	return IsMatch(val, `\w[-._\w]*@\w[-._\w]*\.\w+`)
}

func IsPhone(val string) bool {
	if strings.HasPrefix(val, "+") {
		return IsMatch(val[1:], `\d{13}`)
	} else {
		return IsMatch(val, `\d{11}`)
	}
}

func IsIPv4(val string) bool {
	return IsMatch(val, `((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}`)
}

func IPContains(str string) bool {
	reg := `(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})(\.(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})){3}`
	validatorReg := regexp.MustCompile(reg)
	if len(validatorReg.Find([]byte(str))) == 0 {
		return false
	}
	return true
}
