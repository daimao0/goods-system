package utils

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ToStr 类型转化
func ToStr(e interface{}) string {
	value := reflect.ValueOf(e)
	switch value.Type() {
	case reflect.TypeOf(int(0)), reflect.TypeOf(int8(0)), reflect.TypeOf(int16(0)), reflect.TypeOf(int32(0)), reflect.TypeOf(int64(0)):
		return strconv.FormatInt(value.Int(), 10)
	case reflect.TypeOf(uint(0)), reflect.TypeOf(uint8(0)), reflect.TypeOf(uint16(0)), reflect.TypeOf(uint32(0)), reflect.TypeOf(uint64(0)):
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.TypeOf(float32(0)):
		return strconv.FormatFloat(value.Float(), 'f', -1, 32)
	case reflect.TypeOf(float64(0)):
		return strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.TypeOf(time.Time{}):
		return value.Interface().(time.Time).Format("2006-01-02 15:04:05") // 使用合适的格式化方式
	default:
		return fmt.Sprintf("%v", e) // 对于其他类型，可以采用默认的格式化方法
	}
}

// ToFloat64 类型转化,注意类型转化失败会传默认的0
func ToFloat64(e interface{}) float64 {
	value := reflect.ValueOf(e)
	switch value.Type() {
	case reflect.TypeOf(int(0)), reflect.TypeOf(int8(0)), reflect.TypeOf(int16(0)), reflect.TypeOf(int32(0)), reflect.TypeOf(int64(0)):
		return float64(value.Int())
	case reflect.TypeOf(uint(0)), reflect.TypeOf(uint8(0)), reflect.TypeOf(uint16(0)), reflect.TypeOf(uint32(0)), reflect.TypeOf(uint64(0)):
		return float64(value.Uint())
	case reflect.TypeOf(float32(0)):
	case reflect.TypeOf(float64(0)):
		return value.Float()
	case reflect.TypeOf(""):
		v, _ := strconv.ParseFloat(value.String(), 64)
		return v
	default:
		return math.NaN()
	}
	return math.NaN()
}

// ToFloat32 类型转化,注意类型转化失败会传默认的0
func ToFloat32(e interface{}) float32 {
	value := reflect.ValueOf(e)
	switch value.Type() {
	case reflect.TypeOf(int(0)), reflect.TypeOf(int8(0)), reflect.TypeOf(int16(0)), reflect.TypeOf(int32(0)), reflect.TypeOf(int64(0)):
		return float32(value.Int())
	case reflect.TypeOf(uint(0)), reflect.TypeOf(uint8(0)), reflect.TypeOf(uint16(0)), reflect.TypeOf(uint32(0)), reflect.TypeOf(uint64(0)):
		return float32(value.Uint())
	case reflect.TypeOf(float32(0)):
	case reflect.TypeOf(float64(0)):
		return float32(value.Float())
	case reflect.TypeOf(""):
		v, _ := strconv.ParseFloat(value.String(), 32)
		return float32(v)
	default:
		return float32(math.NaN())
	}
	return float32(math.NaN())
}

// ToInt64 类型转化,注意类型转化失败会传默认的0
func ToInt64(e interface{}) int64 {
	value := reflect.ValueOf(e)
	switch value.Type() {
	case reflect.TypeOf(int(0)), reflect.TypeOf(int8(0)), reflect.TypeOf(int16(0)), reflect.TypeOf(int32(0)), reflect.TypeOf(int64(0)):
		return value.Int()
	case reflect.TypeOf(uint(0)), reflect.TypeOf(uint8(0)), reflect.TypeOf(uint16(0)), reflect.TypeOf(uint32(0)), reflect.TypeOf(uint64(0)):
		return int64(value.Uint())
	case reflect.TypeOf(float32(0)):
	case reflect.TypeOf(float64(0)):
		return int64(math.Round(value.Float()))
	case reflect.TypeOf(""):
		v, _ := strconv.ParseFloat(value.String(), 64)
		return int64(math.Round(v))
	default:
		return int64(math.NaN())
	}
	return int64(math.NaN())
}

// ToInt 类型转化
func ToInt(e interface{}) int {
	value := reflect.ValueOf(e)
	switch value.Type() {
	case reflect.TypeOf(int(0)), reflect.TypeOf(int8(0)), reflect.TypeOf(int16(0)), reflect.TypeOf(int32(0)), reflect.TypeOf(int64(0)):
		return int(value.Int())
	case reflect.TypeOf(uint(0)), reflect.TypeOf(uint8(0)), reflect.TypeOf(uint16(0)), reflect.TypeOf(uint32(0)), reflect.TypeOf(uint64(0)):
		return int(value.Uint())
	case reflect.TypeOf(float32(0)):
	case reflect.TypeOf(float64(0)):
		return int(math.Round(value.Float()))
	case reflect.TypeOf(""):
		v, _ := strconv.ParseFloat(value.String(), 64)
		return int(math.Round(v))
	default:
		return int(math.NaN())
	}
	return int(math.NaN())
}

// ToUInt64 类型转化
func ToUInt64(e interface{}) uint64 {
	value := reflect.ValueOf(e)
	switch value.Type() {
	case reflect.TypeOf(int(0)), reflect.TypeOf(int8(0)), reflect.TypeOf(int16(0)), reflect.TypeOf(int32(0)), reflect.TypeOf(int64(0)):
		return uint64(value.Int())
	case reflect.TypeOf(uint(0)), reflect.TypeOf(uint8(0)), reflect.TypeOf(uint16(0)), reflect.TypeOf(uint32(0)), reflect.TypeOf(uint64(0)):
		return value.Uint()
	case reflect.TypeOf(float32(0)):
	case reflect.TypeOf(float64(0)):
		return uint64(math.Round(value.Float()))
	case reflect.TypeOf(""):
		v, _ := strconv.ParseFloat(value.String(), 64)
		return uint64(math.Round(v))
	default:
		return uint64(math.NaN())
	}
	return uint64(math.NaN())
}

// ToUInt64Arr 类型转化
func ToUInt64Arr(ids string) []uint64 {
	idArr := strings.Split(ids, ",")
	var nums []uint64
	for _, idStr := range idArr {
		nums = append(nums, ToUInt64(idStr))
	}
	return nums
}

// ToStrByUInt64 类型转化
func ToStrByUInt64(ids []uint64) string {
	str := ""
	for _, id := range ids {
		str = str + fmt.Sprintf("%d,", id)
	}
	runes := []rune(str)
	if len(runes) > 0 {
		return string(runes[:len(runes)-1])
	}
	return str
}

// ToTime 类型转化
func ToTime(timeStr string) time.Time {
	if len(timeStr) == len(time.DateOnly) {
		parse, err := time.Parse(time.DateOnly, timeStr)
		if err != nil {
			return time.Time{}
		}
		return parse
	}
	if len(timeStr) == len(time.DateTime) {
		parse, err := time.Parse(time.DateTime, timeStr)
		if err != nil {
			return time.Time{}
		}
		return parse
	}
	return time.Time{}
}
