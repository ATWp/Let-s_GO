package models

import "time"

type OwnAlarm struct {
	ID             int
	WelcomeMessage string
	playDate       time.Time
}

func (a *OwnAlarm) SetTime(time.Time) {
	panic("implement me")
}

func (a *OwnAlarm) SetAlarmSound(string) {
	panic("implement me")
}

func (a *OwnAlarm) PlayAlarm() {
	panic("implement me")
}
