package fn

import (
	"fmt"
	"time"
)

func GetPeriodData(periodType string) (start, end, label string) {
	var now = time.Now()
	var dateFormat = "2006-01-02"
	switch periodType {
	case "day":
		var day = now.Format(dateFormat)
		start = fmt.Sprintf("%s 00:00:00", day)
		end = fmt.Sprintf("%s 23:59:59", day)
		label = now.Format("Jan 02")

	case "week":
		var weekDay = int(now.Weekday()) // Sunday: 0
		var weekStart = now.AddDate(0, 0, -1*weekDay)
		start = fmt.Sprintf("%s 00:00:00", weekStart.Format(dateFormat))
		end = fmt.Sprintf("%s 23:59:59", weekStart.AddDate(0, 0, 6).Format(dateFormat))
		_, wk := now.ISOWeek()
		label = fmt.Sprintf("Week %d", wk)

	case "month":
		var date = now.Day()
		var monthStart = now.AddDate(0, 0, -1*(date-1))
		start = fmt.Sprintf("%s 00:00:00", monthStart.Format(dateFormat))
		end = fmt.Sprintf("%s 23:59:59", monthStart.AddDate(0, 1, 0).Add(time.Second*-86401).Format(dateFormat))
		label = now.Format("Jan")

	case "year":
		var year = now.Format("2006")
		var yearStart, _ = time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s-01-01 00:00:00", year))
		start = yearStart.Format("2006-01-02 15:04:05")
		end = yearStart.AddDate(1, 0, 0).Add(time.Second * -1).Format("2006-01-02 15:04:05")
		label = now.Format("2006")
	}
	return start, end, label
}
