package main

import (
	"fmt"
	"strings"
	"time"
)

//GetNorthAndSouthBoundTrafficEntries
func GetNorthAndSouthBoundTrafficEntries(fileName string) (map[time.Time]TrafficEntries, map[time.Time]TrafficEntries) {
	entries := GetRecordsFromFile(fileName)
	// entries.PrintToFile()
	totalEntries := len(entries)
	fmt.Println("Total No Of entries found", totalEntries)
	northBoundTrafficEntries, southBoundTrafficEntries := entries.toNorthAndSouthBoundTrafficEntries()
	if 0 != len(entries) {
		// fmt.Println("Residuals")
		// fmt.Printf("%+v", entries)
		panic(fmt.Sprintf("*** Residuals Found -  %v ***", len(entries)))
	}
	fmt.Println("North Bound Entries --", len(northBoundTrafficEntries))
	fmt.Println("South Bound Entries --", len(southBoundTrafficEntries))
	isMatched := totalEntries == ((len(southBoundTrafficEntries) * 4) + (len(northBoundTrafficEntries) * 2))
	fmt.Println("Is Matched -- ", isMatched)
	if !isMatched {
		panic("Traffic Analysis failed; Cant match input file records to analysed records.")
	}
	// print(northBoundTrafficEntries, "North Bound")
	// print(southBoundTrafficEntries, "South Bound")
	return TrafficEntries(northBoundTrafficEntries).mapTrafficEntriesByDate(), TrafficEntries(southBoundTrafficEntries).mapTrafficEntriesByDate()
}

func (trafficEntries TrafficEntries) mapTrafficEntriesByDate() map[time.Time]TrafficEntries {
	trafficEntriesMappedByDate := map[time.Time]TrafficEntries{}
	for _, trafficEntry := range trafficEntries {
		entryDate := ClearTimeParamsFromDate(trafficEntry.firstPairOfTyersMarkTime)
		if _, isDateExists := trafficEntriesMappedByDate[entryDate]; !isDateExists {
			trafficEntriesMappedByDate[entryDate] = TrafficEntries{}
		}
		trafficEntriesMappedByDate[entryDate] = append(trafficEntriesMappedByDate[entryDate], trafficEntry)
	}
	return trafficEntriesMappedByDate
}

func (fileEntrys *InputFileTrafficEntrys) toNorthAndSouthBoundTrafficEntries() ([]TrafficEntry, []TrafficEntry) {
	var northBoundTrafficEntries []TrafficEntry
	var southBoundTrafficEntries []TrafficEntry
	emptyTrafficEntryRecord := TrafficEntry{}
	inputFileTrafficEntrys := *fileEntrys
	isPreviousSouthBoundTrafficEntryRecordComplete := true
	for InputFileNoOfRecordsToConsider <= len(inputFileTrafficEntrys) {
		inputFileRecordsForDecidingNorthOrSouthBound := InputFileTrafficEntrys(inputFileTrafficEntrys[0:InputFileNoOfRecordsToConsider])
		northBoundTrafficEntry, southBoundTrafficEntry := inputFileRecordsForDecidingNorthOrSouthBound.decideAndGetNorthOrSouthBoundTrafficEntry()

		if northBoundTrafficEntry != emptyTrafficEntryRecord {
			northBoundTrafficEntry.processNorthBoundTrafficEntry(&northBoundTrafficEntries)
		} else if isPreviousSouthBoundTrafficEntryRecordComplete && southBoundTrafficEntry != emptyTrafficEntryRecord {
			southBoundTrafficEntries = append(southBoundTrafficEntries, southBoundTrafficEntry)
			isPreviousSouthBoundTrafficEntryRecordComplete = false
		} else if southBoundTrafficEntry != emptyTrafficEntryRecord {
			southBoundTrafficEntry.processSouthBoundTrafficEntry(&southBoundTrafficEntries)
			isPreviousSouthBoundTrafficEntryRecordComplete = true
		}

		inputFileTrafficEntrys = inputFileTrafficEntrys[InputFileNoOfRecordsToConsider:]
	}
	*fileEntrys = inputFileTrafficEntrys
	return northBoundTrafficEntries, southBoundTrafficEntries
}

func (inputFileTrafficEntrys InputFileTrafficEntrys) decideAndGetNorthOrSouthBoundTrafficEntry() (TrafficEntry, TrafficEntry) {
	if 2 != len(inputFileTrafficEntrys) {
		panic("Expected 2 values")
	}
	// Can get Two A
	// Can get One A and One B
	southBoundTrafficEntry := TrafficEntry{}
	northBoundTrafficEntry := TrafficEntry{}
	firstRecord := inputFileTrafficEntrys[0]
	secondRecord := inputFileTrafficEntrys[1]
	if strings.EqualFold(secondRecord.recordedOnSide, "B") {
		northBoundTrafficEntry = BuildTrafficEntry(firstRecord.recordedTime, secondRecord.recordedTime)
	} else {
		southBoundTrafficEntry = BuildTrafficEntry(firstRecord.recordedTime, secondRecord.recordedTime)
	}
	return southBoundTrafficEntry, northBoundTrafficEntry
}

func (trafficEntry TrafficEntry) processNorthBoundTrafficEntry(northBoundTrafficEntrysSoFarPtr *[]TrafficEntry) {
	northBoundEntrysSoFar := *northBoundTrafficEntrysSoFarPtr
	trafficEntry.estimatedSpeed = CalculateSpeedInKmph(AverageDistanceFromFirstToSeondAxleOfVehicleInMtrs, trafficEntry.msBetweenFirstAndSecondPairOfTyers)
	trafficEntry.setEstimatedDistanceFromPreviousVehicle(northBoundEntrysSoFar)
	*northBoundTrafficEntrysSoFarPtr = append(northBoundEntrysSoFar, trafficEntry)
}

func (trafficEntry TrafficEntry) processSouthBoundTrafficEntry(southBoundTrafficEntrysSoFarPtr *[]TrafficEntry) {
	southBoundTrafficEntrysSoFar := *southBoundTrafficEntrysSoFarPtr
	inCompleteSouthBoundTrafficEntry := southBoundTrafficEntrysSoFar[len(southBoundTrafficEntrysSoFar)-1:][0]
	southBoundTrafficEntrysSoFar = southBoundTrafficEntrysSoFar[0 : len(southBoundTrafficEntrysSoFar)-1]
	inCompleteSouthBoundTrafficEntry.secondPairOfTyersMarkTime = trafficEntry.firstPairOfTyersMarkTime
	inCompleteSouthBoundTrafficEntry.msBetweenFirstAndSecondPairOfTyers =
		GetDifferenceBtwnDatesInMs(inCompleteSouthBoundTrafficEntry.firstPairOfTyersMarkTime, inCompleteSouthBoundTrafficEntry.secondPairOfTyersMarkTime)
	inCompleteSouthBoundTrafficEntry.estimatedSpeed = CalculateSpeedInKmph(AverageDistanceFromFirstToSeondAxleOfVehicleInMtrs, inCompleteSouthBoundTrafficEntry.msBetweenFirstAndSecondPairOfTyers)
	inCompleteSouthBoundTrafficEntry.setEstimatedDistanceFromPreviousVehicle(southBoundTrafficEntrysSoFar)
	*southBoundTrafficEntrysSoFarPtr = append(southBoundTrafficEntrysSoFar, inCompleteSouthBoundTrafficEntry)
}

func (trafficEntry *TrafficEntry) setEstimatedDistanceFromPreviousVehicle(trafficEntries []TrafficEntry) {
	if len(trafficEntries) > 0 {
		previousTrafficEntry := trafficEntries[len(trafficEntries)-1:][0]
		trafficEntry.msBetweenPreviousVehicleSecondPairToCurrentVechicleFirstPair = GetDifferenceBtwnDatesInMs(previousTrafficEntry.secondPairOfTyersMarkTime, trafficEntry.firstPairOfTyersMarkTime)
		//Current Vehicle speed and time in ms diffreence from previou vehicle
		trafficEntry.estimatedDistFromPreviousVehicle = CalculateDistanceInMeters(trafficEntry.estimatedSpeed, trafficEntry.msBetweenPreviousVehicleSecondPairToCurrentVechicleFirstPair)
	} else {
		trafficEntry.msBetweenPreviousVehicleSecondPairToCurrentVechicleFirstPair = 0
		trafficEntry.estimatedDistFromPreviousVehicle = 0
	}
}

func PrintTrafficEntries(trafficEntries []TrafficEntry, side string) {
	fmt.Println()
	fmt.Printf("<---- %v ---->", side)
	fmt.Println()
	fmt.Println()
	for index, trafficEntry := range trafficEntries {
		fmt.Printf("\n%+v)\nfirstPairOfTyersMarkTime: %+v \nsecondPairOfTyersMarkTime: %+v \nmsBetweenFirstAndSecondPairOfTyers: %+v \nmsBetweenPreviousVehicleSecondPairToCurrentVechicleFirstPair: %+v \nestimatedSpeed: %+v \nestimatedDistFromPreviousVehicle: %+v\n",
			index,
			trafficEntry.firstPairOfTyersMarkTime,
			trafficEntry.secondPairOfTyersMarkTime,
			trafficEntry.msBetweenFirstAndSecondPairOfTyers,
			trafficEntry.msBetweenPreviousVehicleSecondPairToCurrentVechicleFirstPair,
			trafficEntry.estimatedSpeed,
			trafficEntry.estimatedDistFromPreviousVehicle)
	}
}
