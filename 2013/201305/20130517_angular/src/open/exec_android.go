// +build android

package open

import (
	"os/exec"
)

// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-open/
// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-mime/

func open(input string) *exec.Cmd {
	return exec.Command("am start -a android.intent.action.VIEW -n com.android.browser/.BrowserActivity -d ", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command(appName, input)
}
/*
am start -a android.intent.action.VIEW -d file:///sdcard/myweb/index.html -n com.android.browser/.BrowserActivity
*/
