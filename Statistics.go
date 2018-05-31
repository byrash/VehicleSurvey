package main

import (
	"time"
)

// GetStatsInUserSelectedIntervalSlots
func GetStatsInUserSelectedIntervalSlots(entrysPerDayIn5MinSlots map[time.Time][]TrafficEntrySlot, slotInterval int64) map[time.Time]map[int64][]TrafficEntrySlot {
	results := map[time.Time]map[int64][]TrafficEntrySlot{}
	for date, trafficEntrySlots := range entrysPerDayIn5MinSlots {
		// fmt.Println("Working for date ", date.String())
		results[date] = splitIntoIntervalSlots(slotInterval, trafficEntrySlots)
	}
	return results
}

func splitIntoIntervalSlots(slotSize int64, trafficSlotsIn5MinInterval []TrafficEntrySlot) map[int64][]TrafficEntrySlot {
	noOfiterations := (NoOf5MinsIn24Hour / slotSize)
	trafficSlotsInUserSelectedSlotSize := map[int64][]TrafficEntrySlot{}
	for iteration := int64(1); iteration <= noOfiterations; iteration++ {
		for slotIndex, trafficSlot := range trafficSlotsIn5MinInterval {
			if trafficSlot.slotIndex <= (slotSize * iteration) {
				trafficSlotsInUserSelectedSlotSize[iteration] = append(trafficSlotsInUserSelectedSlotSize[iteration], trafficSlot)
			} else {
				trafficSlotsIn5MinInterval = trafficSlotsIn5MinInterval[slotIndex:]
				break
			}
		}
	}
	return trafficSlotsInUserSelectedSlotSize
}

//GenerateStats
func GenerateStats(trafficEntriesMap map[int64][]TrafficEntrySlot) (time.Time, time.Time, float64, float64, int64, float64, float64, int64) {
	peakStartTime, peakEndTime := time.Time{}, time.Time{}
	countOfVehiclesInPeakSlot := int64(0)
	mergedSlotMap := mergeEntries(trafficEntriesMap)
	totalSpeed := float64(0)
	totalDist := float64(0)
	avgDistForSlot := float64(0)
	avgSpeedForSlot := float64(0)
	avgDistForPeakSlot := float64(0)
	avgSpeedForPeakSlot := float64(0)
	totalNoOfVehiclesForSlot := int64(0)
	totalNoOfVehiclesInPeakSlot := int64(0)
	for _, trafficEntrySlot := range mergedSlotMap {
		if trafficEntrySlot.totalNoOfVehicles != 0 {
			totalDist = totalDist + trafficEntrySlot.totalRoughDistanceBetweenCars
			totalSpeed = totalSpeed + trafficEntrySlot.totalRoughSpeedDistribution
			totalNoOfVehiclesForSlot = totalNoOfVehiclesForSlot + trafficEntrySlot.totalNoOfVehicles
			// fmt.Println("trafficEntrySlot.totalNoOfVehicles", trafficEntrySlot.totalNoOfVehicles)
			if trafficEntrySlot.totalNoOfVehicles > countOfVehiclesInPeakSlot {
				countOfVehiclesInPeakSlot = trafficEntrySlot.totalNoOfVehicles
				peakStartTime = trafficEntrySlot.slotStartTime
				peakEndTime = trafficEntrySlot.slotEndTime
				totalNoOfVehiclesInPeakSlot = trafficEntrySlot.totalNoOfVehicles
				avgDistForPeakSlot = trafficEntrySlot.totalRoughDistanceBetweenCars / float64(trafficEntrySlot.totalNoOfVehicles)
				avgSpeedForPeakSlot = trafficEntrySlot.totalRoughSpeedDistribution / float64(trafficEntrySlot.totalNoOfVehicles)
			}

		}
	}

	// fmt.Println("totalNoOfVehiclesInSlot", totalNoOfVehiclesInSlot)
	// fmt.Println("totalDist", totalDist)
	// fmt.Println("totalSpeed", totalSpeed)
	if totalNoOfVehiclesForSlot > 0 {
		avgDistForSlot = totalDist / float64(totalNoOfVehiclesForSlot)
		avgSpeedForSlot = totalSpeed / float64(totalNoOfVehiclesForSlot)
	}
	// fmt.Println("avgDist", avgDist)
	// fmt.Println("avgSpeed", avgSpeed)
	return peakStartTime, peakEndTime, avgDistForPeakSlot, avgSpeedForPeakSlot, totalNoOfVehiclesInPeakSlot, avgDistForSlot, avgSpeedForSlot, totalNoOfVehiclesForSlot
}

func mergeEntries(trafficEntriesInUserSelectedSlotIntervals map[int64][]TrafficEntrySlot) map[int64]TrafficEntrySlot {
	mergedTrafficEntriesWithInSlot := map[int64]TrafficEntrySlot{}
	for index, trafficEntrys := range trafficEntriesInUserSelectedSlotIntervals {
		mergedTrafficEntriesWithInSlot[index] = makeTrafficEntryFromEntrys(trafficEntrys)
	}
	return mergedTrafficEntriesWithInSlot
}

func makeTrafficEntryFromEntrys(trafficEntrysInUserSelectedIntervalToMerge []TrafficEntrySlot) TrafficEntrySlot {
	mergedTrafficEntrySlot := TrafficEntrySlot{}
	totalVehicles := int64(0)
	totalSpeed := float64(0)
	totalDist := float64(0)
	for index, entry := range trafficEntrysInUserSelectedIntervalToMerge {
		if index == 0 {
			mergedTrafficEntrySlot.slotStartTime = entry.slotStartTime
		}
		if index == len(trafficEntrysInUserSelectedIntervalToMerge)-1 {
			mergedTrafficEntrySlot.slotEndTime = entry.slotEndTime
		}
		if entry.totalNoOfVehicles != 0 {
			// fmt.Println("totalNoOfVehicles", entry.totalNoOfVehicles, "roughSpeedDistribution", entry.totalRoughSpeedDistribution, "roughDistanceBetweenCars", entry.totalRoughDistanceBetweenCars)
			totalVehicles = totalVehicles + entry.totalNoOfVehicles
			totalSpeed = totalSpeed + entry.totalRoughSpeedDistribution
			totalDist = totalDist + entry.totalRoughDistanceBetweenCars
		}
	}
	if totalVehicles != 0 {
		mergedTrafficEntrySlot.totalNoOfVehicles = totalVehicles
		// fmt.Println("totalVehicles", totalVehicles)
		// fmt.Println("totalDist", totalDist)
		// fmt.Println("totalSpeed", totalSpeed)
		mergedTrafficEntrySlot.totalRoughDistanceBetweenCars = totalDist
		mergedTrafficEntrySlot.totalRoughSpeedDistribution = totalSpeed
	}

	// for _, trafficEntrySlot := range trafficEntrysInUserSelectedIntervalToMerge {
	// 	str := fmt.Sprintf("slotIndex %v \n slotStartTime %v \n slotEndTime %v \n Total Traffic Entries %v\n roughDistanceBetweenCars %v\n roughSpeedDistribution %v\n", trafficEntrySlot.slotIndex, trafficEntrySlot.slotStartTime.String(), trafficEntrySlot.slotEndTime.String(), len(trafficEntrySlot.trafficEntrys), trafficEntrySlot.roughDistanceBetweenCars, trafficEntrySlot.roughSpeedDistribution)
	// 	fmt.Println(str)
	// }
	// str := fmt.Sprintf("slotIndex %v \n slotStartTime %v \n slotEndTime %v \n Total Traffic Entries %v\n roughDistanceBetweenCars %v\n roughSpeedDistribution %v\n", mergedTrafficEntrySlot.slotIndex, mergedTrafficEntrySlot.slotStartTime.String(), mergedTrafficEntrySlot.slotEndTime.String(), mergedTrafficEntrySlot.totalNoOfVehicles, mergedTrafficEntrySlot.totalRoughDistanceBetweenCars, mergedTrafficEntrySlot.totalRoughSpeedDistribution)
	// fmt.Println(str)
	return mergedTrafficEntrySlot
}
