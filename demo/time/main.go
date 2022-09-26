package main

import (
	"fmt"
	"time"
)

func main()  {
	now, _ := time.Parse("2006-01-02", "2022-01-12")
	//now := time.Now()
	day1 := GetYesterday(now).Format("2006-01-02")
	day2 := GetLastOfWeekDay(now).Format("2006-01-02")
	day3 := GetLastOfMonthDay(now).Format("2006-01-02")
	day4 := GetLastOfQuarterDay(now).Format("2006-01-02")
	day5 := GetLastOfYearDay(now).Format("2006-01-02")

	day6 := GetLastOfWeekDay(now).AddDate(0, 0, -6).Format("2006-01-02")
	day7 := GetFirstOfMonthDay(now).Format("2006-01-02")
	day8 := GetFirstOfQuarterDay(now).Format("2006-01-02")
	day9 := GetLastOfYearDay(now).AddDate(-1, 0, 1).Format("2006-01-02")

	fmt.Println(day1, day2, day3, day4, day5)
	fmt.Println(day6, day7, day8, day9)
}

func GetYesterday(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}

func GetLastOfWeekDay(now time.Time) time.Time {
	// 计算本周周一的偏移值
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	// 计算上周周末的时间
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset - 1)
}

func GetLastOfMonthDay(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, - 1)
}

func GetFirstOfMonthDay(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, -1, 0)
}

func GetLastOfQuarterDay(now time.Time) time.Time {
	var (
		day int
		year = now.Year()
		month = now.Month() - 3
	)
	if month <= 0 {
		day = 31
		month = 12
		year -= 1
	} else if month >= 1 && month <= 3 {
		day = 31
		month = 3
	} else if month >= 4 && month <= 6 {
		day = 30
		month = 6
	} else if month >= 7 && month <= 9 {
		day = 30
		month = 9
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func GetFirstOfQuarterDay(now time.Time) time.Time {
	var (
		day   = 1
		year  = now.Year()
		month = now.Month() - 3
	)
	if month <= 0 {
		month = 10
		year -= 1
	} else if month >= 1 && month <= 3 {
		month = 1
	} else if month >= 4 && month <= 6 {
		month = 4
	} else if month >= 7 && month <= 9 {
		month = 7
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func GetLastOfYearDay(now time.Time) time.Time {
	return time.Date(now.Year() - 1, 12, 31, 0, 0, 0, 0, time.Local)
}