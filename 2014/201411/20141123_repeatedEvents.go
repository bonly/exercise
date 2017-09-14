package main

import (
	"fmt"
	"github.com/PuloV/ics-golang"
)

func main() {
	parser := ics.New()
	ics.FilePath = "tmp/new/"
	ics.DeleteTempFiles = false
	ics.RepeatRuleApply = true
	inputChan := parser.GetInputChan()
	inputChan <- "20141123_repeatedEvents.ics"

	// wait to kill the main goroute
	parser.Wait()
	cal, _ := parser.GetCalendars()
	for _, e := range cal[0].GetEvents() {
		fmt.Printf("%s \n", &e)
		fmt.Println("More info: ", e.GetDescription(), e.GetStart().Format(ics.YmdHis), e.GetSequence())
	}

}