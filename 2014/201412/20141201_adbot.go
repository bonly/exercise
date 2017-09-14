package main

import (
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/cs8425/go-adbbot/adbbot"
)

var verbosity = flag.Int("v", 2, "verbosity")
var ADB = flag.String("adb", "adb", "adb exec path")
var DEV = flag.String("dev", "", "select device")

var APP = flag.String("app", "com.android.vending", "app package name")
var TMPL = flag.String("tmpl", "tmpl.png", "app package name")

func main() {

	log.SetFlags(log.Ldate|log.Ltime)
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	adbbot.Verbosity = *verbosity
	bot := adbbot.NewBot(*DEV, *ADB)


	Vlogln(2, "[adb]", "wait-for-device")
	_, err := bot.Run("wait-for-device")
	if err != nil {
		Vlogln(1, "adb err", err)
	}

	// press Home key
	bot.KeyHome()

	// start APP
	bot.StartApp(*APP)


	// create matching region between Point <100,635> and <9999,9999>
	//reg := bot.NewRectAbs(100, 635, 9999, 9999)

	// or All the screen (slow)
	reg := bot.NewRectAll()

	// create matching template
	tmpl, err := bot.NewTmpl(*TMPL, reg)
	if err != nil {
		Vlogln(2, "load template image err", err)
	} else {

		// try to find target
		// 10 times with 1000ms delay between each search
		x, y, val := bot.FindExistReg(tmpl, 10, 1000)
		if x == -1 && y == -1 {
			Vlogln(2, "template not found", x, y, val)
		} else {
			Vlogln(2, "template found at", x, y, val)
		}

	}

	infoname := time.Now().Format("20060102_150405")
	err = bot.SaveScreen(infoname + ".png")
	if err != nil {
		Vlogln(2, "SaveScreen err", err)
	} else {
		Vlogln(2, "SaveScreen as file ", infoname + ".png")
	}

	// force-stop APP
	bot.KillApp(*APP)

}

func Vlogln(level int, v ...interface{}) {
	if level <= *verbosity {
		log.Println(v...)
	}
}
