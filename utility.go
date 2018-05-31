package main

import (
	"fmt"
	"strings"
	"time"
)

/*
OneKmphToMetersPerSecond holds what i kmoh is equals to in meters per second
*/
const (
	OneKmphToMetersPerSecond                                = 0.28
	AverageDistanceFromFirstToSeondAxleOfVehicleInMtrs      = 2.5
	DefinedSpeedLimitinKmph                                 = 60
	ASensorToBSensorHitAverageTimeInMsShouldBeLessThanValue = 5
	OneSecondToMilliSecond                                  = 1000
)

//http://www.softschools.com/formulas/physics/distance_speed_time_formula/75/
// s = speed (meters/second)
// d = distance traveled (meters)
// t = time (seconds)

// CalculateSpeedInKmph caluclates speed in KMPH for given distance in Meters and time in Milli seconds
func CalculateSpeedInKmph(distanceInMeters float64, timeInMs float64) float64 {
	// s = d / t
	timeInSeconds := MsToSeconds(timeInMs)
	speedInMetersPerSecond := distanceInMeters / timeInSeconds
	return speedInMetersPerSecond / OneKmphToMetersPerSecond
}

//CalculateDistanceInMeters
func CalculateDistanceInMeters(speedInKmph float64, timeInMs float64) float64 {
	// d= s*t
	speedInMetersPerSeond := speedInKmph * OneKmphToMetersPerSecond
	timeInSeconds := MsToSeconds(timeInMs)
	return speedInMetersPerSeond * timeInSeconds
}

//MsToSeconds
func MsToSeconds(timeInMs float64) float64 {
	return timeInMs / OneSecondToMilliSecond
}

/*
AddMilliSecondsToTime will parse the given ms string and will apply that othe date supplied
*/
func AddMilliSecondsToTime(onDate time.Time, timeStrInMs string) (time.Time, error) {
	duration, err := time.ParseDuration(strings.Join([]string{timeStrInMs, "ms"}, ""))
	return onDate.Add(duration), err
}

/*
GetDayStrat Return current time zone start of the day
*/
func GetDayStrat() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

/*
ClearTimeParamsFromDate Return current time zone start of the day
*/
func ClearTimeParamsFromDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

//Add5MinsToTime
func Add5MinsToTime(t time.Time) time.Time {
	// fiveMinDuration, _ := time.ParseDuration("5m")
	// return t.Add(fiveMinDuration)
	return AddMinsToTime(t, 5)
	// return time.Date(t.Year(), t.Month(), t.Day(), 0, (t.Minute() + 5), 0, 0, t.Location())
}

//Add5MinsToTime
func AddMinsToTime(t time.Time, mins int) time.Time {
	minsDuration, _ := time.ParseDuration(fmt.Sprintf("%vm", mins))
	return t.Add(minsDuration)
	// return time.Date(t.Year(), t.Month(), t.Day(), 0, (t.Minute() + 5), 0, 0, t.Location())
}

/*
GetDiffFromGetDifferenceBtwnDatesInMs Previous gets difference between two dates in MS
*/
func GetDifferenceBtwnDatesInMs(previousDate, currentDate time.Time) float64 {
	diff := currentDate.Sub(previousDate)
	return diff.Seconds() * 1000
}

//GetHoursAndBalanceMins
func GetHoursAndBalanceMins(number int64) (int64, int64) {
	number = number * 5
	hours := number / 60
	bal := number % 60
	return hours, bal
}
