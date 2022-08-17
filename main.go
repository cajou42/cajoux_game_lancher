package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Display_laucher(){
	a := app.New()
	w := a.NewWindow("welcome to cajoux laucher")
	var s *fyne.Container
	s = container.New(layout.NewGridLayout(4))
	search_bar := widget.NewEntry()
	search_bar.SetPlaceHolder("search a game")

	s.Add(widget.NewButton("add game", func() {Display_add_game_window(a.NewWindow("add game"))}))
	s.Add(search_bar)
	s.Add(widget.NewButton("search", func() {search()}))
	s.Add(widget.NewButton("custom layout", func() {}))

	Display_game(w,s)
	w.Resize(fyne.NewSize(800,400))
	w.ShowAndRun()
}


func Display_game(w fyne.Window, s *fyne.Container){

		//read game_path.txt
		f, err := os.Open("game_path.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)

		var s2 *fyne.Container
		var box *container.Scroll
		s2 = container.New(layout.NewVBoxLayout())

		//display games layout
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			x := widget.NewButton(scanner.Text(), func() {})
			s2.Add(x)
		}

		box = container.NewVScroll(s2)
		box.SetMinSize(fyne.NewSize(100,400))

		w.SetContent(container.NewVBox(
			s,
			box,
		),)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
}

func search(){

}

func main() {
	Display_laucher()
}