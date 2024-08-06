package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Period struct {
	startTime string
	endTime   string
}

type Schedule struct {
	dayOfWeek   string
	time        Period
	weeks       string
	location    string
	classId     int
	subjectId   string
	subjectName string
	teacherName string
}

func parseTxt(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	timeTable := make([]Schedule, len(lines))
	for i := 0; i < len(timeTable); i++ {
		fields := strings.Split(lines[i], "\t")
		misc := strings.Split(fields[0], ",")
		id, err := strconv.Atoi(fields[3])
		if err != nil {
			log.Fatal(err)
		}
		timeTable[i] = Schedule{
			dayOfWeek: string([]rune(misc[0])[4]),
			time: Period{
				startTime: misc[1][:strings.IndexRune(misc[1], '-')-1],
				endTime:   misc[1][strings.IndexRune(misc[1], '-')+2:],
			},
			weeks:       fields[1],
			location:    fields[2],
			classId:     id,
			subjectId:   fields[6],
			subjectName: fields[7],
			teacherName: fields[10],
		}
		fmt.Println(timeTable[i])
	}
}
