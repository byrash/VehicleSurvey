package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func GetNorthAndSouthBoundEntrysPerDayIn5MinSlots(inputEntrysFileName string) (map[time.Time][]TrafficEntrySlot, map[time.Time][]TrafficEntrySlot) {
	northBoundTrafficEntriesPerDay, southBoundTrafficEntriesPerDay := GetNorthAndSouthBoundTrafficEntries(inputEntrysFileName)
	// for date, trafficEntries := range northBoundTrafficEntriesPerDay {
	// 	fmt.Println(date)
	// 	PrintTrafficEntries([]TrafficEntry(trafficEntries), "North")
	// }
	// for date, trafficEntries := range southBoundTrafficEntriesPerDay {
	// 	fmt.Println(date)
	// 	PrintTrafficEntries([]TrafficEntry(trafficEntries), "South")
	// }
	northBoundTrafficEntrySlotsIn5MinInterval := getTrafficEntriesInSlots(northBoundTrafficEntriesPerDay)
	southBoundTrafficEntrySlotsIn5MinInterval := getTrafficEntriesInSlots(southBoundTrafficEntriesPerDay)
	// printEntriesToFile("north", northBoundTrafficEntrysInSlots)
	// printEntriesToFile("south", southBoundTrafficEntrysInSlots)
	return northBoundTrafficEntrySlotsIn5MinInterval, southBoundTrafficEntrySlotsIn5MinInterval
}

func getTrafficEntriesInSlots(trafficEntriesEachDay map[time.Time]TrafficEntries) map[time.Time][]TrafficEntrySlot {
	trafficEntrySlotsPerDayIN5MinIntervals := map[time.Time][]TrafficEntrySlot{}
	for date, trafficEntries := range trafficEntriesEachDay {
		// fmt.Println(date.String())
		trafficEntrySlotsPerDayIN5MinIntervals[date] = trafficEntries.getTrafficEntrysIn5MinSlots(date)
	}
	return trafficEntrySlotsPerDayIN5MinIntervals
}

func printEntriesToFile(whichBound string, trafficEntrySlotsMap map[time.Time][]TrafficEntrySlot) {
	file, err := os.Create(whichBound + "_slots.txt")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	defer file.Close()
	for date, trafficEntrySlots := range trafficEntrySlotsMap {
		// fmt.Println(date.String())
		printToFile(file, date, trafficEntrySlots)
	}
}

func printToFile(file io.Writer, date time.Time, trafficEntrySlots []TrafficEntrySlot) {
	buffWritter := bufio.NewWriter(file)
	buffWritter.Write([]byte("\n\nFor Date --> " + date.String() + "\n"))
	for _, trafficEntrySlot := range trafficEntrySlots {
		str := fmt.Sprintf("slotIndex %v \n slotStartTime %v \n slotEndTime %v \n Total Traffic Entries %v\n roughDistanceBetweenCars %v\n roughSpeedDistribution %v\n", trafficEntrySlot.slotIndex, trafficEntrySlot.slotStartTime.String(), trafficEntrySlot.slotEndTime.String(), len(trafficEntrySlot.trafficEntrys), trafficEntrySlot.totalRoughDistanceBetweenCars, trafficEntrySlot.totalRoughSpeedDistribution)
		buffWritter.Write([]byte(str))
	}
	buffWritter.Flush()
}

func (trafficEntry TrafficEntry) isTrafficEntryTimeBetweenSlotStartAndEndTime(startTimeInclusive, endTimeExclusive time.Time) bool {
	timeBeingCompared := trafficEntry.firstPairOfTyersMarkTime
	return (timeBeingCompared.Equal(startTimeInclusive) || timeBeingCompared.After(startTimeInclusive)) && timeBeingCompared.Before(endTimeExclusive)
}

func (trafficEntries TrafficEntries) getTrafficEntrysIn5MinSlots(date time.Time) []TrafficEntrySlot {
	trafficEntrySlots := []TrafficEntrySlot{}
	collectTrafficEntrysIn5MinSlots(&trafficEntrySlots, ClearTimeParamsFromDate(date), int64(1), trafficEntries)
	fillMissingSlots(&trafficEntrySlots, date)
	return trafficEntrySlots
}

func fillMissingSlots(trafficEntrySlots *[]TrafficEntrySlot, date time.Time) {
	totalNoOfTrafficEntrySlots := len(*trafficEntrySlots)
	for index := 0; index < NoOf5MinsIn24Hour; index++ {
		// Fill any missing entries
		trafficEntrySlot := TrafficEntrySlot{}
		if index < totalNoOfTrafficEntrySlots {
			trafficEntrySlot = (*trafficEntrySlots)[index]
		}
		// fmt.Println(index, " -> ", trafficEntrySlot.slotIndex)
		if trafficEntrySlot.slotIndex == 0 {
			missingSlotNum := index + 1
			// fmt.Println("Missing Slot ", missingSlotNum)
			slotStartTime := AddMinsToTime(date, (index * 5))
			slotEndTime := Add5MinsToTime(slotStartTime)
			*trafficEntrySlots = append(*trafficEntrySlots, createTrafficEntrySlot(slotStartTime, slotEndTime, int64(missingSlotNum)))
		}
	}
	// return trafficEntrySlots
}

func collectTrafficEntrysIn5MinSlots(trafficEntrySlots *[]TrafficEntrySlot, startTime time.Time, slotIndex int64, trafficEntries TrafficEntries) {
	if len(trafficEntries) <= 0 {
		return
	}
	slotStartTime := startTime
	slotEndTime := Add5MinsToTime(slotStartTime)
	trafficEntrySlot := createTrafficEntrySlot(slotStartTime, slotEndTime, slotIndex)
	var subTrafficEntries TrafficEntries
	totalVehicles := int64(0)
	for index, trafficEntry := range trafficEntries {
		isStratTimeWithInSlot := trafficEntry.isTrafficEntryTimeBetweenSlotStartAndEndTime(slotStartTime, slotEndTime)
		// fmt.Println("Slot Start ", slotStartTime.String(), " Slot End ", slotEndTime.String(), " isStratTimeWithInSlot ", isStratTimeWithInSlot)
		if isStratTimeWithInSlot {
			trafficEntrySlot.trafficEntrys = append(trafficEntrySlot.trafficEntrys, trafficEntry)
			totalVehicles++
		} else {
			subTrafficEntries = trafficEntries[index:]
			break
		}
	}
	trafficEntrySlot.totalNoOfVehicles = totalVehicles
	trafficEntrySlot.calcAndSetTotalSpeedAndDistanceOfSlot()
	*trafficEntrySlots = append(*trafficEntrySlots, trafficEntrySlot)
	slotIndex++
	// fmt.Println("New Slot From ", slotEndTime.String())
	collectTrafficEntrysIn5MinSlots(trafficEntrySlots, slotEndTime, slotIndex, subTrafficEntries)
}

func (trafficEntrySlot *TrafficEntrySlot) calcAndSetTotalSpeedAndDistanceOfSlot() {
	totalItems := len(trafficEntrySlot.trafficEntrys)
	if totalItems <= 0 {
		return
	}
	speedTotal := float64(0.0)
	distTotal := float64(0.0)
	for _, trafficEntry := range trafficEntrySlot.trafficEntrys {
		speedTotal = speedTotal + trafficEntry.estimatedSpeed
		distTotal = distTotal + trafficEntry.estimatedDistFromPreviousVehicle
	}
	trafficEntrySlot.totalRoughDistanceBetweenCars = distTotal
	trafficEntrySlot.totalRoughSpeedDistribution = speedTotal
}

func createTrafficEntrySlot(slotStartTime, slotEndTime time.Time, slotIndex int64) TrafficEntrySlot {
	trafficEntrySlot := TrafficEntrySlot{}
	trafficEntrySlot.slotStartTime = slotStartTime
	trafficEntrySlot.slotEndTime = slotEndTime
	trafficEntrySlot.slotIndex = slotIndex
	return trafficEntrySlot
}
