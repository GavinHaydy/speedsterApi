package utils

import (
	"fmt"
	"log"
	"time"
)

type DateTime time.Time

func (d DateTime) MarshalJSON() ([]byte, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	if err != nil {
		log.Println("解析时间失败:", err)
		return []byte(nil), err
	}
	if time.Time(d) == parsedTime {
		return []byte("null"), err
	}
	dateTime := fmt.Sprintf("%q", time.Time(d).Format("2006-01-02 15:04:05"))
	return []byte(dateTime), nil
}

// String 方法用于提供自定义的字符串表示
func (dt DateTime) String() string {
	return time.Time(dt).Format("2006-01-02 15:04:05")
}
