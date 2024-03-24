package utils

import "time"

func ParseDate(dateStr string) time.Time {
	date, err := time.Parse("06年1月02日", dateStr)
	if err != nil {
		return time.Now()
	}
	return date
}

func GetDayStartDate(move int) time.Time {
	now := time.Now()
	moveDate := now.AddDate(0, 0, move)
	moveStart := time.Date(moveDate.Year(), moveDate.Month(), moveDate.Day(), 0, 0, 0, 0, moveDate.Location())
	return moveStart
}

func GetDayEndDate(move int) time.Time {
	now := time.Now()
	moveDate := now.AddDate(0, 0, move+1)
	endStart := time.Date(moveDate.Year(), moveDate.Month(), moveDate.Day(), 0, 0, 0, 0, moveDate.Location())
	return endStart.Add(-time.Second)
}

func GetDayRangeTime(move int) (start time.Time, end time.Time) {
	start = GetDayStartDate(move)
	end = GetDayEndDate(move)
	return
}

func GetTimeDiffOfDays(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}
