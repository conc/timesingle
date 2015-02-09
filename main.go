package main

import(
	"log"
	"./timesingle"
)

func main() {
	timeSingle := timesingle.NewTTimeSignal()
	timeSingle.WeekDay = 4
	timeSingle.TimeStr = "17:52:10"
	timeSingle.TimeTntervalSecond = 1
	timeSingle.Begin()

	for {
		<-timeSingle.Ch
		log.Println("have one single!")
	}

	return
}
