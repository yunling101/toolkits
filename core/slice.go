package core

import "reflect"

// IntReverse 反转一个整数数组
func IntReverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// StringReverse 反转一个字符串数组
func StringReverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// StringContains 字符串数组是否包含
func StringContains(str []string, k string) bool {
	for _, v := range str {
		if v == k {
			return true
		}
	}
	return false
}

// IntContains 整数数组是否包含
func IntContains(str []int, k int) bool {
	for _, v := range str {
		if v == k {
			return true
		}
	}
	return false
}

// StructContains 通用结构体的某个字段是否包含
func StructContains(obj interface{}, field string, value interface{}) bool {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i).Elem()
			fieldValue := elem.FieldByName(field)
			if fieldValue.IsValid() && fieldValue.Interface() == value {
				return true
			}
		}
	}
	return false
}

// MapContains map某个字段是否包含
func MapContains(obj map[string]interface{}, field string, value interface{}) bool {
	for k, v := range obj {
		if k == field && v == value {
			return true
		}
	}
	return false
}
