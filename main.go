// Gopherjs Vecty Electron's quick start sample

//go:generate gopherjs build main.go -o dist/main.js
//go:generate gopherjs build ./contents -o dist/contents.js
//go:generate cp index.html package.json dist/
//go:generate cp -r css js dist/

package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/gopherjs/gopherjs/js"
)

var (
	mainWindow *js.Object
	require    = js.Global.Get("require")
	electron   = require.Invoke("electron")
)

func main() {
	fmt.Printf("js: %v\n", js.Global.Get("process").Get("version"))
	fmt.Printf("electron: %v\n", js.Global.Get("process").Get("versions").Get("electron"))

	app := electron.Get("app")
	browserWindow := electron.Get("BrowserWindow")
	crashReporter := electron.Get("crashReporter")

	crashReporter.Call("start", map[string]interface{}{
		"companyName": "fake",
		"submitURL":   "localhost",
	})

	app.Call("on", "window-all-closed", func() {
		if js.Global.Get("process").Get("platform").String() != "darwin" {
			app.Call("quit")
		}
	})

	app.Call("on", "ready", func() {
		geom := map[string]int{"width": 800, "height": 600}
		mainWindow = browserWindow.New(geom)

		// and load the index.html of the app.
		_, filename, _, _ := runtime.Caller(1)
		url := "file:///" + path.Join(path.Dir(filename), "index.html")
		mainWindow.Call("loadURL", url)

		// Emitted when the window is closed.
		mainWindow.Call("on", "closed", func() {
			mainWindow = nil
		})
	})
}
