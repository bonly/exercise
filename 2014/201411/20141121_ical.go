package main 

import (
"github.com/laurent22/ical-go/ical"
"fmt"
"time"
)

func main(){
	ny, err := time.LoadLocation("America/New_York");
	if err != nil {
		panic(err);
	}

	createdAt   := time.Date(2010, time.January, 1, 12, 0, 1, 0, time.UTC);
	modifiedAt  := createdAt.Add(time.Second);
	startsAt    := createdAt.Add(time.Second * 2).In(ny);
	endsAt      := createdAt.Add(time.Second * 3).In(ny);

	item := ical.CalendarItem {
		Id: "123",
		CreatedAtUTC: &createdAt,
		ModifiedAtUTC: &modifiedAt,
		StartAt: &startsAt,
		EndAt: &endsAt,
		Summary: "Foo Bar",
		Location: "Berlin\nGermany",
	};

	output := item.Serialize();
	fmt.Printf("%s\n", output);

	node, _ := ical.ParseCalendar(output);

	// val, find := node.DigProperty("SUMMARY");
	val, find := node.DigParameter("DTSTART","20100101T120003Z");
	if find {
		fmt.Printf("%#v\n", val);
	}
	


}
