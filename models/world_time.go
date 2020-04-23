package models

import (
	"errors"
	"fmt"
	"time"
)

type WorldTime struct {
	Id      int
	Hours   int
	Minutes int
	Seconds int
	Day     int
	Month   int
	Year    int
	Tz      string
}

var (
	times []*WorldTime
	/*
		times  []*WorldTime{
			{
				id: 0,
				Hours: 0,
				Minutes: 0,
				Seconds: 0,
				Day: 0,
				Month: 0,
				Year: 1900,
				Tz: "America/Los_Angeles",
			},
			{
				id: 1,
				Hours: 12,
				Minutes: 30,
				Seconds: 0,
				Day: 13,
				Month: 2,
				Year: 2020,
				Tz: "America/Los_Angeles",
			},
		}
		nextid = 2
	*/
	nextid = 1
)

func GetTimeStr() string {
	return time.Now().String()
}

func InitTimeWithTZ(tz string) WorldTime {
	t := time.Now()
	world_time := WorldTime{
		Hours:   t.Hour(),
		Minutes: t.Minute(),
		Seconds: t.Second(),
		Day:     t.Day(),
		Month:   int(t.Month()),
		Year:    t.Year(),
		Tz:      tz,
	}
	return world_time
}

func InitTime() WorldTime {
	return InitTimeWithTZ("America/Los_Angeles")
}

func TimeToStr(t WorldTime) string {
	return fmt.Sprintf("%02d:%02d:%02d %02d/%02d/%4d",
		t.Hours,
		t.Minutes,
		t.Seconds,
		t.Month,
		t.Day,
		t.Year,
	)
}

func GetTimes() []*WorldTime {
	fmt.Println("GetTimes, times =", times)
	return times
}

func AddTime(t WorldTime) (WorldTime, error) {
	fmt.Println("AddTime before, times =", times)
	fmt.Println("and t =", t)
	if t.Id != 0 {
		return WorldTime{}, errors.New("A new WorldTime must not include an id")
	}
	t.Id = nextid
	nextid++
	times = append(times, &t)
	fmt.Println("AddTime after, times =", times)
	return t, nil
}

func GetWorldTimeByTZ(tz string) (WorldTime, error) {
	for _, t := range times {
		if t.Tz == tz {
			return *t, nil
		}
	}
	return WorldTime{}, fmt.Errorf("No WorldTime found with the timezone %v", tz)
}

func GetWorldTimeByID(id int) (WorldTime, error) {
	for _, t := range times {
		if t.Id == id {
			return *t, nil
		}
	}
	return WorldTime{}, fmt.Errorf("No WorldTime found with ID %v", id)
}

func UpdateWorldTime(wt WorldTime, tz string) (WorldTime, error) {
	for i, t := range times {
		if t.Id == wt.Id {
			times[i] = &wt
			return wt, nil
		}
	}
	return WorldTime{}, fmt.Errorf("No WorldTime found with ID %v", wt.Id)
}

func RemoveWorldTime(id int) error {
	for i, t := range times {
		if t.Id == id {
			times = append(times[:i], times[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No WorldTime found with ID %v", id)
}
