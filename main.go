package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	// Open the events.log file
	file, err := os.Open("events.log")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a map to store the NOK events per minute
	eventsPerMinute := make(map[string]int)

	// Create a regular expression to parse the log lines
	regex := regexp.MustCompile(`^\[(.+)\] (.+)$`)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Extract the timestamp and event status from the log line
		match := regex.FindStringSubmatch(line)
		if len(match) != 3 {
			continue
		}
		timestamp, status := match[1], match[2]

		// Parse the timestamp
		t, err := time.Parse("2006-01-02 15:04:05", timestamp)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Convert the timestamp to the nearest minute
		t = t.Truncate(time.Minute)

		// Increment the count of NOK events for the minute
		if status == "NOK" {
			key := t.Format("2006-01-02 15:04")
			eventsPerMinute[key]++
		}
	}

	// Print the count of NOK events per minute
	for k, v := range eventsPerMinute {
		fmt.Printf("%s: %d\n", k, v)
	}
}
