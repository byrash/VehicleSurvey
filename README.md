# This Project is built using golang

## How to run

go run DataType.go utility.go InputFileParser.go TrafficAnalysis.go TrafficEntrySlot.go Statistics.go VehicleSurvey.go CommandLineInput.go

## Problem: Vehicle Survey Code Challenge

```A small city government recently bought a vehicle counter. In order for the vehicle counter to work, pneumatic rubber hoses are stretched across the road. Data is produced by the vehicle counter as traffic drives across the hoses. The city government requires a program to interpret the data that the machine produces.
The data from the machine looks like this:
! A268981
! A269123
! A604957
! B604960Vehicle Survey Code Challenge
A small city government recently bought a vehicle counter. In order for the vehicle counter to work, pneumatic rubber hoses are stretched across the road. Data is produced by the vehicle counter as traffic drives across the hoses. The city government requires a program to interpret the data that the machine produces.
The data from the machine looks like this:
! A268981
! A269123
! A604957
! B604960
! A605128
! B605132
! A1089807
! B1089810
! A1089948
! B1089951
The numbers are the number of milliseconds since midnight when the mark occurred. Thus, the first line above represents a pair of tires driving by at 12:04:28am. The second line represents another pair of tires going by 142ms later (almost certainly the 2nd axle of the car).
The vehicle counter has two pneumatic rubber hoses - one stretches across both lanes of traffic, and one goes just across traffic in one direction. Each hose independently records when tires drive over it. As such, cars going in one direction (say, northbound) only record on one sensor (preceded with an 'A'), while cars going in the direction (say, southbound) are recorded on both sensors. Lines 3-6 above represent a second car going in the other direction. The first set of tires hit the 'A' sensor at 12:10:04am, and then hit the 'B' sensor 3ms later. The second set of tires then hit the 'A' sensor 171ms later, and then the 'B' sensor 4ms later.
The machine was left to run for 5 days in a row (starting on a Monday). This is obvious because the times in the data make several sudden drops:
! A86328771
! B86328774
! A86328899
! B86328902
! A582668
! B582671
! A582787
! B582789
The city has asked you to see how many analysis features you can provide:
 • Total vehicle counts in each direction: morning versus evening, per hour, per half hour, per 20 minutes, and per 15 minutes.
• The above counts can be displayed for each day of the session, or you can see averages across all the days.
• Peak volume times.
• The (rough) speed distribution of traffic.
• Rough distance between cars during various periods.

### Luckily for you, you know that:

• The speed limit on the road where this is recorded is 60kph (that doesn't mean that everyone drives this speed, or that no one exceeds it, but it's a good starting point).
• The average wheelbase of cars in the city is around 2.5 metres between axles.
• Only 2-axle vehicles were allowed on this road during the recording sessions.
This coding challenge should be accompanied with a data file containing the full vehicle survey data to be analysed by the program.Vehicle Survey Code Challenge
```

## __Algorithm__

1. Read the input file and convert the entries from file into which side ( A or B / North or South Bound) and what time ( Converting from MS to DateTime ). Date will be starting from system date. (InputFileParser.go)
2. Covert all Entries parsed from file to North Bound and South Bound Entries along with estimating speed and distance from previous vehicle (TrafficAnalysis.go)
     1. Algorithm for analyzing which side or bound the input file records corresponds to is as follows
     2. Pick 2 records from input file parsed
     3. if both the records has A side ( North bound) make those records as North bound
     4. If one the record has B side ( South bound )
         1. if we already holding one partial B entry merge that entry to this entry and make it as South Bound Entry
         2. If no hold of partial B entry exits, hold the current entry as partial entry and continue with next records, Step 2.2
3. Generate all traffic Entries into __5 Minute slot intervals ( As 5 min intervals will cover all scenarios, 15 (5*3), 20 (5*4), ... )__, while holding total no of vehicles with in the slot, total speed, total distance from previous vehicle and entries too
4. Take User Input from command line and convert it into Slot Size. i.e if User Selected 20 Mins, for 20 mins we will have 4, 5 minutes Slots. So the user selected interval Slot mapped from 20 mins to system interval of 4.(Rem we are basing our default interval as 5 Min slot from Step 3) (CommandLineInput.go)
5. On User user valid input parse the input data file and generate Statistics and dump to file (VehicleSurvey.go and Statistics.go)

### Sample Execution

```
$ go run DataType.go utility.go InputFileParser.go TrafficAnalysis.go TrafficEntrySlot.go Statistics.go VehicleSurvey.go CommandLineInput.go
Total No Of entries found 67296
North Bound Entries -- 11096
South Bound Entries -- 11276
Is Matched --  true
1) Morning versus Evening
2) Per hour
3) Per half hour
4) Per 20 minutes
5) Per 15 minutes
Enter your choice:( 1 or 2 or 3 or 4 or 5) or Ctrl + C to exit
1


<----------- North Bound -------------->
For Date --> 2018-05-31 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-05-31 12:00:00 +1000 AEST
First Peak End Time: 2018-06-01 00:00:00 +1000 AEST
Rough Distance for peak time -- 512.3067103025262 Meters
Rough Speed for peak time -- 63.12268446220949 KMPH
Total No Of Vehciles in Peak Slot -- 1486
----Stats for day-----
Rough Distance for day -- 705.8779756784186 Meters
Rough Speed for day -- 63.00019858418278 KMPH
Total Vehicles for day -- 2190

For Date --> 2018-06-01 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-06-01 12:00:00 +1000 AEST
First Peak End Time: 2018-06-02 00:00:00 +1000 AEST
Rough Distance for peak time -- 511.2526704957544 Meters
Rough Speed for peak time -- 63.18942096317052 KMPH
Total No Of Vehciles in Peak Slot -- 1467
----Stats for day-----
Rough Distance for day -- 690.595825013355 Meters
Rough Speed for day -- 63.056832567442065 KMPH
Total Vehicles for day -- 2194

For Date --> 2018-06-02 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-06-02 12:00:00 +1000 AEST
First Peak End Time: 2018-06-03 00:00:00 +1000 AEST
Rough Distance for peak time -- 509.9729152533388 Meters
Rough Speed for peak time -- 63.12387476731166 KMPH
Total No Of Vehciles in Peak Slot -- 1489
----Stats for day-----
Rough Distance for day -- 684.8617323983842 Meters
Rough Speed for day -- 63.18304885194596 KMPH
Total Vehicles for day -- 2216

For Date --> 2018-06-03 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-06-03 12:00:00 +1000 AEST
First Peak End Time: 2018-06-04 00:00:00 +1000 AEST
Rough Distance for peak time -- 495.54910652456005 Meters
Rough Speed for peak time -- 62.9377484858406 KMPH
Total No Of Vehciles in Peak Slot -- 1513
----Stats for day-----
Rough Distance for day -- 669.1297629714975 Meters
Rough Speed for day -- 62.83898531061633 KMPH
Total Vehicles for day -- 2275

For Date --> 2018-06-04 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-06-04 12:00:00 +1000 AEST
First Peak End Time: 2018-06-05 00:00:00 +1000 AEST
Rough Distance for peak time -- 513.1244499880108 Meters
Rough Speed for peak time -- 63.03605424006325 KMPH
Total No Of Vehciles in Peak Slot -- 1481
----Stats for day-----
Rough Distance for day -- 693.7248792478017 Meters
Rough Speed for day -- 62.98299366624465 KMPH
Total Vehicles for day -- 2221

<----------- South Bound -------------->
For Date --> 2018-05-31 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-05-31 12:00:00 +1000 AEST
First Peak End Time: 2018-06-01 00:00:00 +1000 AEST
Rough Distance for peak time -- 645.4413950462408 Meters
Rough Speed for peak time -- 63.14019232264554 KMPH
Total No Of Vehciles in Peak Slot -- 1170
----Stats for day-----
Rough Distance for day -- 679.7958274853283 Meters
Rough Speed for day -- 62.99716998414593 KMPH
Total Vehicles for day -- 2224

For Date --> 2018-06-01 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-06-01 12:00:00 +1000 AEST
First Peak End Time: 2018-06-02 00:00:00 +1000 AEST
Rough Distance for peak time -- 634.540171586241 Meters
Rough Speed for peak time -- 63.107107304211596 KMPH
Total No Of Vehciles in Peak Slot -- 1192
----Stats for day-----
Rough Distance for day -- 653.5279615938178 Meters
Rough Speed for day -- 62.92500610084075 KMPH
Total Vehicles for day -- 2333

For Date --> 2018-06-02 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-06-02 12:00:00 +1000 AEST
First Peak End Time: 2018-06-03 00:00:00 +1000 AEST
Rough Distance for peak time -- 685.0346647724048 Meters
Rough Speed for peak time -- 62.761037197803965 KMPH
Total No Of Vehciles in Peak Slot -- 1115
----Stats for day-----
Rough Distance for day -- 687.0946224016167 Meters
Rough Speed for day -- 62.92110203583511 KMPH
Total Vehicles for day -- 2220

For Date --> 2018-06-03 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-06-03 12:00:00 +1000 AEST
First Peak End Time: 2018-06-04 00:00:00 +1000 AEST
Rough Distance for peak time -- 665.067784553828 Meters
Rough Speed for peak time -- 62.88339894830245 KMPH
Total No Of Vehciles in Peak Slot -- 1123
----Stats for day-----
Rough Distance for day -- 669.7760263125567 Meters
Rough Speed for day -- 62.91408420944529 KMPH
Total Vehicles for day -- 2241

For Date --> 2018-06-04 00:00:00 +1000 AEST
----Peak-----
First Peak Start Time: 2018-06-04 12:00:00 +1000 AEST
First Peak End Time: 2018-06-05 00:00:00 +1000 AEST
Rough Distance for peak time -- 647.4782142008232 Meters
Rough Speed for peak time -- 62.96698638376691 KMPH
Total No Of Vehciles in Peak Slot -- 1172
----Stats for day-----
Rough Distance for day -- 672.6983084752497 Meters
Rough Speed for day -- 62.97936332605248 KMPH
Total Vehicles for day -- 2258

1) Morning versus Evening
2) Per hour
3) Per half hour
4) Per 20 minutes
5) Per 15 minutes
Enter your choice:( 1 or 2 or 3 or 4 or 5) or Ctrl + C to exit
```