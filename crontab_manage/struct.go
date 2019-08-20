package crontabmanage

// Crontab : Crontabの設定状況
type Crontab struct {
	Month      int
	DayOfMonth int
	Day        int
	Hour       int
	Minute     int
	User       string
}
