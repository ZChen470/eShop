package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Vector []float32

// 将 Vector 转为数据库中的值 pgvector 的格式
// pgvector 允许通过字符串的形式写入向量
func (v Vector) Value() (driver.Value, error) {
	strs := make([]string, len(v))
	for i, f := range v {
		strs[i] = fmt.Sprintf("%f", f)
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ",")), nil
}

// 从数据库中读取 pgvector 的值
func (v *Vector) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("Vector: cannot convert %T to string", str)
	}
	str = strings.Trim(str, "[]")
	parts := strings.Split(str, ",")
	vec := make([]float32, len(parts))

	for i, part := range parts {
		var val float32
		_, err := fmt.Sscanf(strings.TrimSpace(part), "%f", &val)
		if err != nil {
			return err
		}
		vec[i] = val
	}
	*v = vec
	return nil
}
