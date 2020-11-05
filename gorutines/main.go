package main

import "GOLANG_FAST/models"

func main() {
	ownAlarm := models.ownAlarm{
		ID:             1,
		WelcomeMessage: "Hello",
	}
	alarmExample(ownAlarm)
}

func alarmExample(alarm models.Alarm) {

}
