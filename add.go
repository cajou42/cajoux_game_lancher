package main

import (
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Display_add_game_window(w fyne.Window){
	fmt.Println("coucou")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter game binary path")

	w.SetContent(container.NewVBox(
		input, 
		widget.NewButton("add a game", func() {
		write_game_path(input.Text)
		input.Text = ""
	}),))
	w.Show()
}

func write_game_path(txt string){
	path := []byte(txt)
	err := ioutil.WriteFile("game_path.txt", path, 0644)
	if err != nil {
        fmt.Println(err)
    }
}