package timex

import (
	"time"
	_ "time/tzdata"
)

var shanghai *time.Location

func init() {
	var err error
	shanghai, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

// TimeFormat 默认格式
const TimeFormat = "2006-01-02 15:04:05"

// Format 转上海时区并格式化
func Format(t time.Time) string {
	return t.In(shanghai).Format(TimeFormat)
}

// FormatPtr 格式化 *time.Time
func FormatPtr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.In(shanghai).Format(TimeFormat)
}

// FormatWithLayout 自定义格式
func FormatWithLayout(t time.Time, layout string) string {
	return t.In(shanghai).Format(layout)
}

// Now 获取上海当前时间
func Now() time.Time {
	return time.Now().In(shanghai)
}
