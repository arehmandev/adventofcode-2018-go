package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

const timeFormat = "[2006-01-02 15:04]"

type logEntry struct {
	timestamp time.Time
	message   string
}

type logEntryList []logEntry

// Forward request for length
func (p logEntryList) Len() int {
	return len(p)
}

// Define compare
func (p logEntryList) Less(i, j int) bool {
	return p[i].timestamp.Before(p[j].timestamp)
}

// Define swap over an array
func (p logEntryList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type guard struct {
	id              int
	logs            logEntryList
	maxSleepMin     int
	maxSleepMinFreq int
	totalSleep      int
}

type guardlist []guard

// Forward request for length
func (p guardlist) Len() int {
	return len(p)
}

// Define compare
func (p guardlist) Less(i, j int) bool {
	return p[i].maxSleepMinFreq < p[j].maxSleepMinFreq
}

// Define swap over an array
func (p guardlist) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	timelog := strings.Split(string(data), "\n")

	// Sort logs by timestamp
	events := logEntryList{}
	for _, entry := range timelog {
		entryMessage := strings.Join(strings.Split(string(entry), " ")[2:], " ")
		entryRawTimestamp := strings.Join(strings.Split(string(entry), " ")[0:2], " ")
		entryTimestamp, _ := time.Parse(timeFormat, entryRawTimestamp)

		newentry := logEntry{timestamp: entryTimestamp, message: entryMessage}
		events = append(events, newentry)
	}
	sort.Sort(events)

	// Create a map of guard -> logs
	var currentGuardID int
	guardMap := make(map[int]guard)

	for _, entry := range events {
		if strings.Contains(entry.message, "Guard") {
			currentGuardIDString := strings.Replace(strings.Split(entry.message, " ")[1], "#", "", -1)
			currentGuardID, _ = strconv.Atoi(currentGuardIDString)
			continue
		}

		newguard := getOrCreate(currentGuardID, guardMap)
		newguard.logs = append(newguard.logs, entry)
		guardMap[currentGuardID] = newguard
	}

	// Loop through guards and find maxslept value
	guardList := guardlist{}

	for _, guardObj := range guardMap {
		var startTime time.Time
		var endTime time.Time

		minuteMap := make(map[int]int)
		for _, entry := range guardObj.logs {

			if strings.Contains(entry.message, "falls asleep") {
				startTime = entry.timestamp
			} else {
				endTime = entry.timestamp

				totalslept := int(endTime.Sub(startTime).Minutes())
				guardObj.totalSleep += totalslept

				for index := startTime.Minute(); index < endTime.Minute(); index++ {
					minuteMap[index]++
					if minuteMap[index] > guardObj.maxSleepMinFreq {
						guardObj.maxSleepMin = index
						guardObj.maxSleepMinFreq = minuteMap[index]
					}
				}

			}
		}
		guardList = append(guardList, guardObj)
	}

	sort.Sort(guardList)

	highestGuard := guardList[len(guardList)-1]
	fmt.Println(highestGuard.id, highestGuard.maxSleepMin, highestGuard.maxSleepMinFreq, highestGuard.id*highestGuard.maxSleepMin)
}

func getOrCreate(guardID int, guardMap map[int]guard) guard {

	if currentGuard, ok := guardMap[guardID]; ok {
		return currentGuard
	}

	guardMap[guardID] = guard{id: guardID}
	return guardMap[guardID]

}
