package comic

import "time"

type I interface {
	Name() string
	URLOf(date time.Time) string
}
