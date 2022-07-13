package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Display_add_game_window(w fyne.Window){

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter game binary path")

	w.SetContent(container.NewVBox(
		input, 
		widget.NewButton("add a game", func() {
		write_game_path(input.Text)
		input.Text = ""
	}),))
	w.Resize(fyne.NewSize(800,400))
	w.Show()
}

func write_game_path(txt string){
	path := []byte(txt + "\n")
	// err := ioutil.WriteFile("game_path.txt", path, 0644)
	// if err != nil {
    //     fmt.Println(err)
    // }
	f, err := os.OpenFile("game_path.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(string(path)); err != nil {
		panic(err)
	}
}