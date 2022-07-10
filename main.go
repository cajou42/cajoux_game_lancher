package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Display_laucher(){
	fmt.Println("alu")
	a := app.New()
	w := a.NewWindow("welcome to cajoux laucher")
	search_bar := widget.NewEntry()
	search_bar.SetPlaceHolder("search a game")

	w.SetContent(container.NewVBox(
		widget.NewButton("add game", func() {
			Display_add_game_window(a.NewWindow("add game"))
		}),
		search_bar,
		widget.NewButton("search", func() {
			search()
		}),
		widget.NewButton("custom layout", func() {
			fmt.Println("alu")
		}),
	))

	w.ShowAndRun()
}

func search(){

}

func main() {
	Display_laucher()
}