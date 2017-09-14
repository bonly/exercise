package main

import (
	"fmt"
	"github.com/PuloV/ics-golang"
	"time"
)

func main() {

	//  create new parser
	parser := ics.New()

	// get the input chan
	inputChan := parser.GetInputChan()

	// send the calendar urls to be parsed
	inputChan <- "http://www.google.com/calendar/ical/bg.bulgarian%23holiday%40group.v.calendar.google.com/public/basic.ics"
	inputChan <- "http://www.google.com/calendar/ical/en.bulgarian%23holiday%40group.v.calendar.google.com/public/basic.ics"

	//  wait for the calendar to be parsed
	parser.Wait()

	// get all calendars in this parser
	cal, err := parser.GetCalendars()

	//  check for errors
	if err == nil {
		// New Years Eve is holiday !
		t, _ := time.Parse(ics.YmdHis, "2016-03-27 00:00:00")
		for _, calendar := range cal {
			// print calendar info
			fmt.Println(calendar.GetName(), calendar.GetDesc())

			//  get the events for the New Years Eve
			eventsForDay, _ := calendar.GetEventsByDate(t)

			for i, event := range eventsForDay {
				// print event info (event implements Stringer interface)
				fmt.Printf("№%d  %s \n", i, event)
			}
		}

	}

}