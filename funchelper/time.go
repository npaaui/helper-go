package funchelper

import "time"

/****************************
 * 时间相关
 ****************************/
func GetNow() (date string) {
	date = time.Now().Format("2006-01-02 15:04:05")
	return
}

func GetForever() (date string) {
	date = "2099-01-01 00:00:00"
	return
}

func GetBegin() (date string) {
	date = "2001-01-01 00:00:00"
	return
}

func GetTomorrowBegin() (date string) {
	date = time.Now().AddDate(0, 0, 1).Format("2006-01-02 00:00:00")
	return
}

func GetAfterHour(hour int) (date string) {
	date = time.Now().Add(time.Hour * time.Duration(hour)).Format("2006-01-02 15:04:05")
	return
}
