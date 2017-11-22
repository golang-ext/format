package format

import (
	"strconv"
	"fmt"
)

/**
 * 格式化字符串
 */
func Format(i interface{}, format ...interface{}) string {
	if len(format) == 0 {
		if i == nil {
			return ""
		}
		switch i.(type) {
		case int:
			return fmt.Sprintf("%d", i)
		case float64:
			return fmt.Sprintf("%f", i)
		case float32:
			return fmt.Sprintf("%f", i)
		}
	}
	return fmt.Sprintf(format[0].(string), append(format[1:], i)...)
}

/**
 * TODO: int => string
 * @param i int数字
 * @return string
 */
func Itoa(i int) string {
	return strconv.Itoa(i)
}

func FormatInt8(i int8, base int) string {
	return strconv.FormatInt(int64(i), base)
}
func FormatInt16(i int16, base int) string {
	return strconv.FormatInt(int64(i), base)
}
func FormatInt32(i int32, base int) string {
	return strconv.FormatInt(int64(i), base)
}
func FormatInt64(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

func FormatUInt8(i uint8, base int) string {
	return strconv.FormatUint(uint64(i), base)
}
func FormatUInt16(i uint16, base int) string {
	return strconv.FormatUint(uint64(i), base)
}
func FormatUInt32(i uint32, base int) string {
	return strconv.FormatUint(uint64(i), base)
}
func FormatUInt64(i uint64, base int) string {
	return strconv.FormatUint(i, base)
}

// FormatFloat 将浮点数 f 转换为字符串值
// f：要转换的浮点数
// fmt：格式标记（b、e、E、f、g、G）
// prec：精度（数字部分的长度，不包括指数部分）
// bitSize：指定浮点类型（32:float32、64:float64）
//
// 格式标记：
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
//
// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
func FormatFloat32(f float32, prec int) string {
	return strconv.FormatFloat(float64(f), 'f', prec, 32)
}
func FormatFloat64(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

func FormatBool(b bool) string {
	return strconv.FormatBool(b)
}
