package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_DATE_FILE = "sample.txt"
)

func main() {
	northBoundPerDayIn5MinSlots, southBoundPerDayIn5MinSlots := GetNorthAndSouthBoundEntrysPerDayIn5MinSlots(INPUT_DATE_FILE)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("1) Morning versus Evening")
		fmt.Println("2) Per hour")
		fmt.Println("3) Per half hour")
		fmt.Println("4) Per 20 minutes")
		fmt.Println("5) Per 15 minutes")
		fmt.Println("Enter your choice:( 1 or 2 or 3 or 4 or 5) or Ctrl + C to exit")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		num, err := strconv.ParseInt(text, 10, 0)
		if err != nil {
			fmt.Println("\n--Invalid Input--[", text, "]")
			fmt.Println(err)
		} else {
			interval := getIntervalToSlotSize(num)
			if interval == 0 {
				fmt.Println("Invalid Input. Try again..!!!")
				continue
			}
			// fmt.Println("Interval being Considered -- ", interval)
			fmt.Println()
			fmt.Println()
			VechicleSurvey(interval, northBoundPerDayIn5MinSlots, "North Bound")
			VechicleSurvey(interval, southBoundPerDayIn5MinSlots, "South Bound")
		}
	}
}

func getIntervalToSlotSize(interval int64) int64 {
	switch interval {
	case 1:
		return NoOf5MinsIn12Hour
	case 2:
		return NoOf5MinsIn1Hour
	case 3:
		return NoOf5MinsIn30Mins
	case 4:
		return NoOf5MinsIn20Mins
	case 5:
		return NoOf5MinsIn15Mins
	default:
		return 0
	}
}
