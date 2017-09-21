package main

import "github.com/gopherjs/vecty"

func main() {
	vecty.SetTitle("Gopherjs Vecty Electron")
	vecty.RenderBody(&PageView{})
}
