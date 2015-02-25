package timesingle

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

/**************************************************
WeekDay : 周几，如果是-1表示每天在TimeStr时间点开始返回时间信号，否则标识星期几开始返回时间信号（注：以7表示星期天）
TimeStr : 定时时间点，格式如15:09:21 hh:mm:ss格式
TimeTntervalSecond : 1*60*60 间隔多长时间返回信号
Ch chan int64： 传递信号的channel
***************************************************/
type TimeSignal struct {
	WeekDay            int64
	TimeStr            string
	TimeTntervalSecond int64
	Ch                 chan int
}

func NewTimeSignal() *TimeSignal {
	var result TimeSignal
	result.Ch = make(chan int)
	return &result
}

func (p *TimeSignal) Begin() {
	go p.beginDeal()
}

func (p *TimeSignal) beginDeal() {
	diffSecond, err := getDiffSecond(p.TimeStr)
	if err != nil {
		panic(err)
		return
	}

	if p.WeekDay == -1 {
		if diffSecond < 0 {
			diffSecond += 24 * 60 * 60
		}
	} else if (p.WeekDay > 0) && (p.WeekDay < 8) {
		diffWeek := p.WeekDay - getWeedOfToday()
		if diffWeek < 0 {
			diffWeek += 7
		}
		diffSecond += diffWeek * 24 * 60 * 60
		if diffSecond < 0 {
			diffSecond += 7 * 24 * 60 * 60
		}
	} else {
		panic(errors.New("error time weedday"))
	}

	time.Sleep(time.Duration(diffSecond) * time.Second)
	p.Ch <- 1
	for {
		time.Sleep(time.Duration(p.TimeTntervalSecond) * time.Second)
		p.Ch <- 1
	}

	return
}

/**************************************************
函数名称：getDiffSecond
函数功能：返回目标时间串到现在时间点间隔的秒数
备    注：可能为负
***************************************************/
func getDiffSecond(timeStr string) (int64, error) {
	nowTime := time.Now().Format(time.RFC3339)
	nowDaySecond, err := getSecondOfToday(nowTime[11:19])
	if err != nil {
		return 0, err
	}

	aimDaySecond, err := getSecondOfToday(timeStr)
	if err != nil {
		return 0, err
	}

	diffSecond := aimDaySecond - nowDaySecond

	return diffSecond, nil
}

/**************************************************
函数名称：getSecondOfToday
函数功能：返回timeStr代表的时间距今天0点的秒数
***************************************************/
func getSecondOfToday(timeStr string) (int64, error) {
	spitArray := strings.Split(timeStr, ":")
	if len(spitArray) != 3 {
		return 0, errors.New("error time string")
	}

	aimHour, err := strconv.ParseInt(spitArray[0], 10, 64)
	if err != nil || aimHour < 0 || aimHour > 23 {
		return 0, errors.New("error time string")
	}
	aimMinute, err := strconv.ParseInt(spitArray[1], 10, 64)
	if err != nil || aimMinute < 0 || aimMinute > 59 {
		return 0, errors.New("error time string")
	}
	aimSecond, err := strconv.ParseInt(spitArray[2], 10, 64)
	if err != nil || aimSecond < 0 || aimSecond > 59 {
		return 0, errors.New("error time string")
	}
	aim := aimHour*3600 + aimMinute*60 + aimSecond

	return aim, nil
}

/**************************************************
函数名称：getWeedOfToday
函数功能：返回今天是周几
备    注：星期天 : 7
***************************************************/
func getWeedOfToday() int64 {
	days := make(map[string]int64)
	days["Sunday"] = 7
	days["Monday"] = 1
	days["Tuesday"] = 2
	days["Wednesday"] = 3
	days["Thursday"] = 4
	days["Friday"] = 5
	days["Saturday"] = 6

	t := time.Now().Weekday().String()
	return days[t]
}
