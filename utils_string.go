package wxxx

import "strings"

func Trim(str *string) string {
	if str == nil {
		return ""
	}
	return strings.Trim(*str, " ")
}

func TrimStr(str string) string {
	return strings.Trim(str, " ")
}

func IsEmpty(str *string) bool {
	return str == nil || len(strings.Trim(*str, " ")) == 0
}
func IsEmptyStr(str string) bool {
	return len(strings.Trim(str, " ")) == 0
}

//noinspection ALL
func IsNotEmpty(str *string) bool {
	return !IsEmpty(str)
}

//noinspection ALL
func IsNotEmptyStr(str string) bool {
	return !IsEmptyStr(str)
}

//noinspection ALL
func IsStrEqual(str1 string, str2 string) bool {
	if IsEmptyStr(str1) && IsEmptyStr(str2) {
		return true
	}
	return str1 == str2
}

//noinspection ALL
func IsStrNotEqual(str1 string, str2 string) bool {
	return !IsStrEqual(str1, str2)
}

//noinspection ALL
func IsEqual(str1 *string, str2 *string) bool {
	if IsEmpty(str1) && IsEmpty(str2) {
		return true
	}
	return *str1 == *str2
}

//noinspection ALL
func IsNotEqual(str1 *string, str2 *string) bool {
	return !IsEqual(str1, str2)
}

//noinspection ALL
func IsEqualStrIgnoreCase(str1 string, str2 string) bool {
	if IsEmptyStr(str1) && IsEmptyStr(str2) {
		return true
	}
	return strings.ToLower(TrimStr(str1)) == strings.ToLower(TrimStr(str2))
}

//noinspection ALL
func ToUpper(str string) string {
	return strings.ToUpper(TrimStr(str))
}
