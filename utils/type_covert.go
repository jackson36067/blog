package utils

import "strconv"

// StringToUint 将string类型的数据转换成uint类型(常用于id的转换)
func StringToUint(s string) (uint, error) {
	id64, err := strconv.ParseUint(s, 10, 64)
	return uint(id64), err
}
