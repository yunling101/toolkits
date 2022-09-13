package date

import (
	"fmt"
	"time"
)

type Date struct {
	Now time.Time
}

// New
func New() *Date {
	return &Date{Now: time.Now()}
}

// 当前多少小时之前时间
func (d *Date) BeforeToHour(before int) time.Time {
	return d.Now.Add(-time.Hour * time.Duration(before))
}

// 当前多少小时之后时间
func (d *Date) AfterToHour(after int) time.Time {
	return d.Now.Add(time.Hour * time.Duration(after))
}

// 当前多少分钟之前时间
func (d *Date) BeforeToMinute(before int) time.Time {
	return d.Now.Add(-time.Minute * time.Duration(before))
}

// 当前多少分钟之后时间
func (d *Date) AfterToMinute(after int) time.Time {
	return d.Now.Add(time.Minute * time.Duration(after))
}

// 天开始时间
func (d *Date) DayStart(n int) time.Time {
	return time.Date(
		d.Now.Year(),
		d.Now.Month(),
		d.Now.Day(), 0, 0, 0, 0, time.Local,
	).AddDate(0, 0, n)
}

// 天结束时间
func (d *Date) DayEnd(n int) time.Time {
	return time.Date(
		d.Now.Year(),
		d.Now.Month(),
		d.Now.Day(), 23, 59, 59, 0, time.Local,
	).AddDate(0, 0, n)
}

// 结束时间
func (d Date) EndTime(n time.Time) time.Time {
	return time.Date(
		n.Year(),
		n.Month(),
		n.Day(), 23, 59, 59, 0, time.Local,
	)
}

// 周第一天
func (d *Date) FirstDateOfWeek() time.Time {
	offset := int(time.Monday - d.Now.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(
		d.Now.Year(),
		d.Now.Month(),
		d.Now.Day(), 0, 0, 0, 0, time.Local,
	).AddDate(0, 0, offset)
}

// 周最后一天
func (d Date) LastDateOfWeek() time.Time {
	offset := int(time.Saturday - d.Now.Weekday())
	if offset == 6 {
		offset = 0
	} else {
		offset += 1
	}

	return time.Date(
		d.Now.Year(),
		d.Now.Month(),
		d.Now.Day(), 23, 59, 59, 0, time.Local,
	).AddDate(0, 0, offset)
}

// 近7天第一天
func (d *Date) BeforeWeekOfFirst() time.Time {
	return d.FirstDateOfWeek().AddDate(0, 0, -7)
}

// 上周最后一天
func (d *Date) AfterWeekOfLast() time.Time {
	s := d.FirstDateOfWeek()
	return time.Date(
		s.Year(),
		s.Month(),
		s.Day(), 23, 59, 59, 0, time.Local,
	).AddDate(0, 0, -1)
}

// 月第一天
func (d Date) FirstDateOfMonth() time.Time {
	dd := d.Now.AddDate(0, 0, -d.Now.Day()+1)
	return time.Date(
		dd.Year(),
		dd.Month(),
		dd.Day(),
		0, 0, 0, 0, dd.Location())
}

// 月最后一天
func (d Date) LastDateOfMonth() time.Time {
	dd := d.FirstDateOfMonth().AddDate(0, 1, -1)
	return time.Date(
		dd.Year(),
		dd.Month(),
		dd.Day(),
		23, 59, 59, 0, dd.Location())
}

// 年第一天
func (d Date) FirstDateOfYear() time.Time {
	return time.Date(
		d.Now.Year(),
		time.January, 1, 0, 0, 0, 0, d.Now.Location())
}

// 年最后一天
func (d Date) LastDateOfYear() time.Time {
	return time.Date(
		d.Now.Year(),
		time.December, 31, 23, 59, 59, 0, d.Now.Location())
}

// 拿前时间当编号
func (d Date) CurrentTimeToCid() string {
	return d.Now.Format("20060102150405")
}

// 当前月份最后一天时间
//func (d *Date) CurrentMonthEnd() time.Time {
//    return time.Date(
//        d.Now.Year(),
//        d.Now.Month(),
//        d.Now.Day(),
//        d.Now.Hour(),
//        d.Now.Minute(),
//        0, 0, d.Now.Location(),
//    )
//}

// 字符月份到时间
func (d *Date) StringMonthToTime(month string) (time.Time, error) {
	m := fmt.Sprintf("%s-01 00:00:00", month)
	return time.ParseInLocation("2006-01-02 15:04:05", m, time.Local)
}

// 字符串转时间
func (d *Date) StringToDate(n, format string) (time.Time, error) {
	return time.ParseInLocation(format, n, time.Local)
}

// 年月日时间到月份格式
func (d *Date) TimeToMonthString(dd time.Time) string {
	return dd.Format("2006-01")
}
