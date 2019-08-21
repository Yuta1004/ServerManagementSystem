package crontabmanage

// Crontab : Crontabの設定状況
type Crontab struct {
	Month      CrontabDateNum
	DayOfMonth CrontabDateNum
	Day        CrontabDateNum
	Hour       CrontabDateNum
	Minute     CrontabDateNum
	User       string
	Command    []string
}

// CrontabDateNum : Crontab表記に合わせた日時要素情報をもつ
type CrontabDateNum struct {
	Wildcard bool
	Nums     []int
	RawData  string
}
