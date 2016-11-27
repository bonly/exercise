package main

import (
    "io/ioutil"
    "log"

    "github.com/google/gxui"
    "github.com/google/gxui/drivers/gl"
    "github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
    theme := dark.CreateTheme(driver)

    window := theme.CreateWindow(1024, 800, "Hi")
    window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray50))

    fontData, err := ioutil.ReadFile("/home/bonly/.fonts/YaHei/YaHei.Consolas.1.11b.ttf") //font comes from windows
    if err != nil {
        log.Fatalf("error reading font: %v", err)
    }
    font, err := driver.CreateFont(fontData, 512)
    if err != nil {
        panic(err)
    }
    label := theme.CreateLabel()
    label.SetFont(font)
    label.SetText("辰蜃")

    window.AddChild(label)

    window.OnClose(driver.Terminate)
}

func main() {
    gl.StartDriver(appMain)
}

/*

diff --git a/v3.1/glfw/glfw/src/x11_window.c b/v3.1/glfw/glfw/src/x11_window.c
index 4f2538b..85656da 100644
--- a/v3.1/glfw/glfw/src/x11_window.c
+++ b/v3.1/glfw/glfw/src/x11_window.c
@@ -913,11 +913,11 @@ static void processEvent(XEvent *event)
                 Status status;
                 wchar_t buffer[16];
 
-                if (XFilterEvent(event, None))
-                {
+//                if (XFilterEvent(event, None))
+//                {
                     // Discard intermediary (dead key) events for character input
-                    break;
-                }
+//                    break;
+//                }
 
                 const int count = XwcLookupString(window->x11.ic,
                                                   &event->xkey,
@@ -1749,6 +1749,8 @@ void _glfwPlatformPollEvents(void)
     {
         XEvent event;
         XNextEvent(_glfw.x11.display, &event);
+	if (XFilterEvent(&event, None))
+		continue;
         processEvent(&event);
     }

*/

