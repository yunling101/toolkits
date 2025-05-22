package time

import (
	"time"
)

const (
	FormatMonth           = "2006-01"
	FormatDate            = "2006-01-02"
	FormatDateTime        = "2006-01-02 15:04:05"
	FormatTime            = "15:04"
	FormatTimeSecond      = "15:04:05"
	FormatTimeMillisecond = "2006-01-02 15:04:05.000"
	FormatTimeStrand      = "20060102150405"
)

// Format 格式化当前时间
func (c *cli) Format(layout string) string {
	return c.now.Format(layout)
}

// StringToTime 字符串转时间
func (c *cli) StringToTime(str, layout string) (time.Time, error) {
	return time.ParseInLocation(layout, str, time.Local)
}
