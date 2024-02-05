package util

import "time"

func GenerateDateList(startDate, endDate string) ([]string, error) {
	// 解析开始日期和结束日期
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, err
	}

	// 初始化日期列表
	var dateList []string

	// 逐天生成日期
	currentDate := start
	for currentDate.Before(end) || currentDate.Equal(end) {
		dateList = append(dateList, currentDate.Format("2006-01-02"))
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dateList, nil
}

func GenerateDayNum(startDate, endDate string) int {
	// 解析开始日期和结束日期
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return 0
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return 0
	}

	return int(end.Sub(start).Hours() / 24)
}

func DiffNumDay(dayDate string, diff int) string {
	day, err := time.Parse("2006-01-02", dayDate)
	if err != nil {
		return dayDate
	}

	return day.AddDate(0, 0, diff).Format(time.DateOnly)
}
