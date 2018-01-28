package b
import (
	_ "unsafe"
	_ "a"
)
//go:linkname say a.say
func say(name string) string
func Greet(name string) string {
	return say(name)
}
func Hi(name string) string
