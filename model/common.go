package model

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/goccy/go-json"
)

// 字符串Json数组
type StrArr []string

func (c StrArr) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
func (c *StrArr) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, c)
}

type DbHookError error
