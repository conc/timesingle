timesingle
==== 
#How to use:

go get github.com/conc/timesingle

#Import:

import github.com/conc/timesingle

#Examples:

```golang
package main

import(
	"log"
	"github.com/conc/timesingle"
)

func main() {
	timeSingle := timesingle.NewTimeSignal()
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
```
