package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Display_laucher(){
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
		}),
	))

	Display_game(w)
	w.Resize(fyne.NewSize(800,400))
	w.ShowAndRun()
}

func Display_game(w fyne.Window){
	s := container.NewVBox(
		widget.NewButton("custom layout", func() {
		}),
	) 
	w.SetContent(container.NewVScroll(
		s,
	))

}

func search(){

}

func main() {
	Display_laucher()
}