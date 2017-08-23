package Figo

import (
	"github.com/quexer/utee"
	"time"
)

func T_SubDays(fromTime, toTime time.Time) int {
	if toTime.Location().String() != fromTime.Location().String() {
		return -1
	}
	if hours := toTime.Sub(fromTime).Hours(); hours <= 0 {
		return -1
	} else if hours < 24 {
		t1y, t1m, t1d := toTime.Date()
		t2y, t2m, t2d := fromTime.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)
		if isSameDay {
			return 0
		} else {
			return 1
		}
	} else {
		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times
			return int(hours / 24)
		} else { // more than 24 hours
			return int(hours/24) + 1
		}
	}
}

func T_EachDayInt(fromTime, toTime time.Time) []int {
	count := T_SubDays(fromTime, toTime)
	times := make([]int, 0)
	t := fromTime
	v := T_DayInt(t)
	times = append(times, v)
	for i := 0; i < count; i++ {
		t = t.Add(time.Hour * 24)
		times = append(times, T_DayInt(t))
	}
	return times
}

func T_DayInt(t time.Time) int {
	v, err := TpInt(t.Format("20060102"))
	utee.Chk(err)
	return v
}
