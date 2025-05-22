package time

import (
	"time"
)

type cli struct {
	now time.Time
}

func New() *cli {
	return &cli{now: time.Now()}
}

// BeforeToHour 当前多少小时之前时间
func (c *cli) BeforeToHour(before int) time.Time {
	return c.now.Add(-time.Hour * time.Duration(before))
}

// AfterToHour 当前多少小时之后时间
func (c *cli) AfterToHour(after int) time.Time {
	return c.now.Add(time.Hour * time.Duration(after))
}

// BeforeToMinute 当前多少分钟之前时间
func (c *cli) BeforeToMinute(before int) time.Time {
	return c.now.Add(-time.Minute * time.Duration(before))
}

// AfterToMinute 当前多少分钟之后时间
func (c *cli) AfterToMinute(after int) time.Time {
	return c.now.Add(time.Minute * time.Duration(after))
}

// DayStartTime 当前几天的开始时间
func (c *cli) DayStartTime(n int) time.Time {
	return time.Date(
		c.now.Year(),
		c.now.Month(),
		c.now.Day(), 0, 0, 0, 0, time.Local,
	).AddDate(0, 0, n)
}

// DayEndTime 当前几天的结束时间
func (c *cli) DayEndTime(n int) time.Time {
	return time.Date(
		c.now.Year(),
		c.now.Month(),
		c.now.Day(), 23, 59, 59, 0, time.Local,
	).AddDate(0, 0, n)
}

// EndTime 结束时间
func (c *cli) EndTime(n time.Time) time.Time {
	return time.Date(
		n.Year(),
		n.Month(),
		n.Day(), 23, 59, 59, 0, time.Local,
	)
}

// FirstDateOfWeek 当前周第一天
func (c *cli) FirstDateOfWeek() time.Time {
	offset := int(time.Monday - c.now.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(
		c.now.Year(),
		c.now.Month(),
		c.now.Day(), 0, 0, 0, 0, time.Local,
	).AddDate(0, 0, offset)
}

// LastDateOfWeek 当前周最后一天
func (c *cli) LastDateOfWeek() time.Time {
	offset := int(time.Saturday - c.now.Weekday())
	if offset == 6 {
		offset = 0
	} else {
		offset += 1
	}

	return time.Date(
		c.now.Year(),
		c.now.Month(),
		c.now.Day(), 23, 59, 59, 0, time.Local,
	).AddDate(0, 0, offset)
}

// BeforeWeekOfFirst 近7天第一天
func (c *cli) BeforeWeekOfFirst() time.Time {
	return c.FirstDateOfWeek().AddDate(0, 0, -7)
}

// AfterWeekOfLast 上周最后一天
func (c *cli) AfterWeekOfLast() time.Time {
	s := c.FirstDateOfWeek()
	return time.Date(
		s.Year(),
		s.Month(),
		s.Day(), 23, 59, 59, 0, time.Local,
	).AddDate(0, 0, -1)
}

// FirstDateOfMonth 当前月第一天
func (c *cli) FirstDateOfMonth() time.Time {
	dd := c.now.AddDate(0, 0, -c.now.Day()+1)
	return time.Date(
		dd.Year(),
		dd.Month(),
		dd.Day(),
		0, 0, 0, 0, dd.Location())
}

// LastDateOfMonth 当前月最后一天
func (c *cli) LastDateOfMonth() time.Time {
	d := c.FirstDateOfMonth().AddDate(0, 1, -1)
	return time.Date(
		d.Year(),
		d.Month(),
		d.Day(),
		23, 59, 59, 0, d.Location())
}

// FirstDateOfYear 当前年第一天
func (c *cli) FirstDateOfYear() time.Time {
	return time.Date(
		c.now.Year(),
		time.January, 1, 0, 0, 0, 0, c.now.Location())
}

// LastDateOfYear 当前年最后一天
func (c *cli) LastDateOfYear() time.Time {
	return time.Date(
		c.now.Year(),
		time.December, 31, 23, 59, 59, 0, c.now.Location())
}
