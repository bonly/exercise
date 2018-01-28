package main

import (
	"grac/comic"
	"time"
)

type xkcd struct {
}

func (xkcd) Name() string {
	return "XKCD"
}

func (xkcd) URLOf(date time.Time) string {
	panic("Not implemented.")
}

func Init() comic.I {
	return &xkcd{}
}
