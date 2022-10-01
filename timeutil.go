package timeutil

import (
	"errors"
	"time"
)

const HoursMinutes = "15:04"
const HoursMinutesSeconds = "15:04:05"

type HMDailyTime string

func NewHMDailyTime(s string) (HMDailyTime, error) {
	if ValidTimeFormat(HoursMinutes, s) {
		return HMDailyTime(s), nil
	}
	return HMDailyTime(""), errors.New("not a valid time in format 15:04")
}

func (h HMDailyTime) GetTime() time.Time {
	r, _ := time.Parse(HoursMinutes, string(h))
	return r
}

func (h HMDailyTime) Valid() bool {
	return ValidTimeFormat(HoursMinutes, string(h))
}

func InTimeRangeDaily(start, end, test time.Time) bool {
	startT := OnlyTime(start)
	endT := OnlyTime(end)
	testT := OnlyTime(test)
	if endT.Before(startT) {
		endT = endT.Add(time.Hour * 24)
	}
	if startT.Before(testT) && endT.After(testT) {
		return true
	}
	testT = testT.Add(time.Hour * 24)
	if startT.Before(testT) && endT.After(testT) {
		return true
	}
	return false
}

func OnlyTime(t time.Time) time.Time {
	t1, _ := time.Parse(HoursMinutesSeconds, t.Format(HoursMinutesSeconds))
	return t1
}

func OnlyDateAndTMZ(t time.Time) time.Time {
	t2, _ := time.Parse("2006-01-02 MST", t.Format("2006-01-02 MST"))
	return t2
}

func CombineDateAndTime(targetDateTmz, targetTime time.Time) time.Time {
	targetDateTmz = OnlyDateAndTMZ(targetDateTmz)
	targetTime = OnlyTime(targetTime)
	return targetDateTmz.Add(time.Hour*time.Duration(targetTime.Hour()) + time.Minute*time.Duration(targetTime.Minute()) + time.Second*time.Duration(targetTime.Second()))
}

func ValidTimeFormat(f, s string) bool {
	_, err := time.Parse(f, s)
	return err == nil
}

func NextOccurenceOf(currentTime, nextTime time.Time) time.Time {
	next := CombineDateAndTime(currentTime, nextTime)
	if currentTime.After(next) {
		next = next.Add(time.Hour * 24)
	}
	return next
}

func ClosestTime(currentTime time.Time, checkTimes []time.Time) int {
	closestDiff := time.Hour * 24 * 30
	closestI := 0
	for ti := range checkTimes {
		diff := checkTimes[ti].Sub(currentTime)
		if diff < 0 {
			diff = -diff
		}
		if diff < closestDiff {
			closestDiff = diff
			closestI = ti
		}
	}
	return closestI
}

func ClosestTimeAfter(currentTime time.Time, checkTimes []time.Time) int {
	closestDiff := time.Hour * 24 * 30
	closestI := 0
	for ti := range checkTimes {
		diff := checkTimes[ti].Sub(currentTime)
		if diff < closestDiff && diff >= 0 {
			closestDiff = diff
			closestI = ti
		}
	}
	return closestI
}

func ClosestTimeBefore(currentTime time.Time, checkTimes []time.Time) int {
	closestDiff := time.Hour * 24 * 30
	closestI := 0
	for ti := range checkTimes {
		diff := checkTimes[ti].Sub(currentTime)
		if diff < closestDiff && diff <= 0 {
			closestDiff = diff
			closestI = ti
		}
	}
	return closestI
}
