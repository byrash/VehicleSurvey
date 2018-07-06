# run: build
# 	./CommandLineInput

DEPS = DataType.go utility.go InputFileParser.go TrafficAnalysis.go TrafficEntrySlot.go Statistics.go VehicleSurvey.go

run: build
	go run CommandLineInput.go $(DEPS)

build: CommandLineInput.go $(DEPS)
	CGO_ENABLED=0 GOOS=linux go build $<

$(DEPS):
	CGO_ENABLED=0 GOOS=linux go build $@