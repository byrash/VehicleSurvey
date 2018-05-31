package main

import (
	"math"
	"testing"
)

func TestCalculateSpeedInKmph(t *testing.T) {
	speedInKmph := CalculateSpeedInKmph(2.5, 148.8095238095)
	actual := math.Round(speedInKmph)
	if actual != 60 {
		t.Errorf("Expected %v but found %v", 60, actual)
	}
}

func TestCalculateDistanceInMeters(t *testing.T) {
	distanceInMtrs := CalculateDistanceInMeters(60, 148.8095238095)
	actual := math.Round(distanceInMtrs)
	if actual != 2 {
		t.Errorf("Expected %v but found %v", 2, actual)
	}
}

func TestAddMilliSecondsToTime(t *testing.T) {
	actual, err := AddMilliSecondsToTime(GetDayStrat(), "1000")
	if err != nil {
		t.Error("Expected to have no errors")
	}
	hrs, mins, sec := actual.Clock()
	if sec != 1 || hrs != 0 || mins != 0 {
		t.Errorf("Expected Only 1 second to be added")
	}
}

func TestGetDayStrat(t *testing.T) {
	dayStart := GetDayStrat()
	hrs, mins, sec := dayStart.Clock()
	if sec != 0 || hrs != 0 || mins != 0 {
		t.Errorf("Expected to be 12 AM")
	}
}

func TestGetDifferenceBtwnDatesInMs(t *testing.T) {
	actual, _ := AddMilliSecondsToTime(GetDayStrat(), "1000")
	ms := GetDifferenceBtwnDatesInMs(GetDayStrat(), actual)
	if ms != 1000 {
		t.Errorf("Expected to be 1000 ms")
	}
}

func TestGetHoursAndBalance5Mins(t *testing.T) {
	hrs, bal := GetHoursAndBalanceMins(3)
	if hrs != 0 || bal != 15 {
		t.Errorf("Expected to be 0 Hrs 15 mins")
	}
	hrs, bal = GetHoursAndBalanceMins(6)
	if hrs != 0 || bal != 30 {
		t.Errorf("Expected to be 0 Hrs 630 mins")
	}
	hrs, bal = GetHoursAndBalanceMins(12)
	if hrs != 1 || bal != 0 {
		t.Errorf("Expected to be 1 Hrs 0 mins")
	}
	hrs, bal = GetHoursAndBalanceMins(13)
	if hrs != 1 || bal != 5 {
		t.Errorf("Expected to be 1 Hrs 5 mins")
	}
	hrs, bal = GetHoursAndBalanceMins(145)
	if hrs != 12 || bal != 5 {
		t.Errorf("Expected to be 12 Hrs 5 mins")
	}
}
