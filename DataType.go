package main

import (
	"time"
)

const (
	NoOf5MinsIn15Mins              = 3
	NoOf5MinsIn20Mins              = 4
	NoOf5MinsIn30Mins              = 6
	NoOf5MinsIn1Hour               = 12
	NoOf5MinsIn12Hour              = 144
	NoOf5MinsIn24Hour              = 288
	InputFileNoOfRecordsToConsider = 2
)

type InputFileTrafficEntrys []InputFileEntry

/*
InputFileEntry structure to load the string parsed to readble date format
*/
type InputFileEntry struct {
	recordedOnSide string
	recordedTime   time.Time
}

// TrafficEntry
type TrafficEntry struct {
	firstPairOfTyersMarkTime                                     time.Time
	secondPairOfTyersMarkTime                                    time.Time
	msBetweenFirstAndSecondPairOfTyers                           float64
	msBetweenPreviousVehicleSecondPairToCurrentVechicleFirstPair float64
	estimatedSpeed                                               float64
	estimatedDistFromPreviousVehicle                             float64
}

// TrafficEntrySlot
type TrafficEntrySlot struct {
	slotIndex                     int64
	slotStartTime                 time.Time
	slotEndTime                   time.Time
	trafficEntrys                 []TrafficEntry
	totalNoOfVehicles             int64
	totalRoughSpeedDistribution   float64
	totalRoughDistanceBetweenCars float64
}

type TrafficEntries []TrafficEntry

/*
BuildTrafficEntry
*/
func BuildTrafficEntry(firstPairOfTyersMarkTime, secondPairOfTyersMarkTime time.Time) TrafficEntry {
	trafficEntry := TrafficEntry{}
	trafficEntry.firstPairOfTyersMarkTime = firstPairOfTyersMarkTime
	trafficEntry.secondPairOfTyersMarkTime = secondPairOfTyersMarkTime
	trafficEntry.msBetweenFirstAndSecondPairOfTyers = GetDifferenceBtwnDatesInMs(firstPairOfTyersMarkTime, secondPairOfTyersMarkTime)
	return trafficEntry
}
