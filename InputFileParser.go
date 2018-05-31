package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
GetRecordsFromFile parses given file and maps to aour custome DS
*/
func GetRecordsFromFile(fileName string) InputFileTrafficEntrys {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to Open file")
		os.Exit(1)
	}
	defer file.Close()
	bufferedScanner := bufio.NewScanner(file)
	if bufferedScanner.Err() != nil {
		log.Fatal(bufferedScanner.Err())
		os.Exit(1)
	}
	var entries InputFileTrafficEntrys
	dayStart := GetDayStrat()
	previousNum := "0"
	for bufferedScanner.Scan() {
		line := bufferedScanner.Text()
		split := strings.Split(line, "")
		currentNum := strings.Join(split[1:], "")
		if isNextDayVal(previousNum, currentNum) {
			// Next day
			twentyFourehrsDuration, _ := time.ParseDuration("24h")
			dayStart = dayStart.Add(twentyFourehrsDuration)
		}
		whenDt, _ := AddMilliSecondsToTime(dayStart, currentNum)
		vehicleInfo := InputFileEntry{recordedOnSide: split[0], recordedTime: whenDt}
		entries = append(entries, vehicleInfo)
		previousNum = currentNum
	}
	return entries
}

/*
PrintToFile prints all records to file for debugging
*/
func (r InputFileTrafficEntrys) PrintToFile() {
	file, _ := os.Create("out.txt")
	buffWritter := bufio.NewWriter(file)
	defer file.Close()
	for _, item := range r {
		str := "Side -> " + item.recordedOnSide + " At -> " + item.recordedTime.String() + "\n"
		buffWritter.Write([]byte(str))
	}
	buffWritter.Flush()
}

func isNextDayVal(previous, current string) bool {
	previousNum, _ := strconv.ParseInt(previous, 10, 64)
	currentNum, _ := strconv.ParseInt(current, 10, 64)
	return currentNum < previousNum
}
