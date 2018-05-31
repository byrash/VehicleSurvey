package main

import (
	"fmt"
	"sort"
	"time"
)

func VechicleSurvey(systemSlotInterval int64, entriesPerDayIn5MinSlots map[time.Time][]TrafficEntrySlot, whichBound string) {
	generateStatsAndPrint(entriesPerDayIn5MinSlots, systemSlotInterval, whichBound)
}

func generateStatsAndPrint(entrysPerDayIn5MinSlots map[time.Time][]TrafficEntrySlot, userSelectedInterval int64, whichBound string) {
	results := GetStatsInUserSelectedIntervalSlots(entrysPerDayIn5MinSlots, userSelectedInterval)
	fmt.Println("<-----------", whichBound, "-------------->")
	for _, date := range getSortedKeysForDate(results) {
		entries := results[date]
		fmt.Println("For Date -->", date.String())
		peakStartTime, peakEndTime, avgDistForPeakSlot, avgSpeedForPeakSlot, totalNoOfVehiclesInPeakSlot, avgDistForDay, avgSpeedForDay, totalNoOfVehiclesForDay := GenerateStats(entries)
		fmt.Println("----Peak-----")
		fmt.Println("First Peak Start Time:", peakStartTime.String(), "\nFirst Peak End Time:", peakEndTime.String())
		fmt.Println("Rough Distance for peak time --", avgDistForPeakSlot, "Meters")
		fmt.Println("Rough Speed for peak time --", avgSpeedForPeakSlot, "KMPH")
		fmt.Println("Total No Of Vehciles in Peak Slot --", totalNoOfVehiclesInPeakSlot)
		fmt.Println("----Stats for day-----")
		fmt.Println("Rough Distance for day --", avgDistForDay, "Meters")
		fmt.Println("Rough Speed for day --", avgSpeedForDay, "KMPH")
		fmt.Println("Total Vehicles for day --", totalNoOfVehiclesForDay)
		fmt.Println("")
		// keysInNaturalOrder := getSortedKeys(entries)
		// for _, key := range keysInNaturalOrder {
		// 	trafficSlotEntrys := entries[int64(key)]
		// 	fmt.Println(key, len(trafficSlotEntrys))
		// }
	}
}

func getSortedKeysForDate(entries map[time.Time]map[int64][]TrafficEntrySlot) []time.Time {
	var keys []time.Time
	for k := range entries {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i].Before(keys[j]) })
	return keys
}

func getSortedKeys(entries map[int64][]TrafficEntrySlot) []int {
	var keys []int
	for k := range entries {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	return keys
}
