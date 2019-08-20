package crontabmanage

// AllocCrontabStruct : Crontab構造体を新しく作って返す
func AllocCrontabStruct() *Crontab {
	c := Crontab{}
	return &c
}
