package util

import (
	"strconv"
	"time"
)

func GetTomorrow0Second() int64 {
	return UTCWest8TodayZeroTime() + 24*time.Hour.Nanoseconds()/1e6
}
func GetTomorrow1Second() int64 {
	return UTCWest8TodayZeroTime() + 25*time.Hour.Nanoseconds()/1e6
}
func GetTomorrow0SecondStamp() time.Time {
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	currentUtcTime := time.Now().In(cstZone)
	timePoint := time.Date(currentUtcTime.Year(), currentUtcTime.Month(), currentUtcTime.Day(), 24, 0, 0, 0, cstZone)
	return timePoint
}
func GetUTCNowStamp() time.Time {
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	t := time.Now().In(cstZone)
	return t
}
func GetYesterday0Second() int64 {
	return UTCWest8TodayZeroTime() - 24*time.Hour.Nanoseconds()/1e6
}

// 获取当前日期字符串 yyyy-mm-dd
func GetNowDateStr() string {
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	unixTimestamp := UTCWest8TodayZeroTime() / 1e3
	dateS := time.Unix(unixTimestamp, 0).In(cstZone).Format("2006-01-02")
	return dateS
}

// 时间戳转字符串日期，到小时， yyyy-mm-dd-HH
func TimeStampToDateHourStr(timestamp int64) string {
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	//日期转化为时间戳
	timeLayout := "2006010215" //转化所需模板
	//时间戳转化为日期
	datetime := time.Unix((timestamp)/1e3, 0).In(cstZone).Format(timeLayout)
	return datetime
}

func GetBeforeDateStr() string {
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	unixTimestamp := (UTCWest8TodayZeroTime() - 24*time.Hour.Nanoseconds()/1e6) / 1e3
	dateS := time.Unix(unixTimestamp, 0).In(cstZone).Format("2006-01-02")
	return dateS
}

func GetDateNow() int64 {
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	t := time.Now().In(cstZone)
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), cstZone)
	return tm1.UnixNano() / 1e6
}
func GetTimePerNHours(n int64) []int64 {
	var timeArray []int64
	for i := 0; i < int(24/n)+1; i++ {
		timeArray = append(timeArray, UTCWest8TodayZeroTime()+int64(i)*n*time.Hour.Nanoseconds()/1e6)
	}
	return timeArray
}

// utc-8 0点时间
func UTCWest8TodayZeroTime() int64 {
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	currentUtcTime := time.Now().In(cstZone)
	timePoint := time.Date(currentUtcTime.Year(), currentUtcTime.Month(), currentUtcTime.Day(), 0, 0, 0, 0, cstZone)
	return timePoint.UnixNano() / 1e6
}
func UTCWest8Date() int32 {
	unixTimestamp := UTCWest8TodayZeroTime() / 1e3
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	dateS := time.Unix(unixTimestamp, 0).In(cstZone).Format("20060102")
	date, _ := strconv.Atoi(dateS)
	return int32(date)
}
func UTCWest8Yesterday() int64 {
	unixTimestamp := GetYesterday0Second() / 1e3
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	dateS := time.Unix(unixTimestamp, 0).In(cstZone).Format("20060102")
	date, _ := strconv.Atoi(dateS)
	return int64(date)
}

func CurrentTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}
func PeriodToTimeStamp(period string) int64 {
	var cstZone = time.FixedZone("utc-8", -8*3600) // 西八
	year, _ := strconv.Atoi(period[0:4])
	month, _ := strconv.Atoi(period[4:6])
	day, _ := strconv.Atoi(period[6:8])
	timePoint := time.Date(year, time.Month(month), day, 0, 0, 0, 0, cstZone)
	return timePoint.UnixNano() / 1e6
}
