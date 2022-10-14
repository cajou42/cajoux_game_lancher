package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Data struct {
	Data []Game `json:"datas"`
}

type Game struct {
	GameName string `json:"GameName"`
	Path string `json:"Path"`
}

func Display_add_game_window(w fyne.Window){

	game_name := widget.NewEntry()
	game_name.SetPlaceHolder("Enter game name")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter game binary path")

	w.SetContent(container.NewVBox(
		input,
		game_name,
		widget.NewButton("add a game", func() {
		write_game_path(game_name.Text, input.Text)
		input.Text = ""
		game_name.Text = ""
	}),))
	w.Resize(fyne.NewSize(800,400))
	w.Show()
}

func write_game_path(name string, txt string){
	// path := []byte(txt + "\n")
	// f, err := os.OpenFile("game_path.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// if _, err = f.WriteString(string(path)); err != nil {
	// 	panic(err)
	// }

	file, _ := ioutil.ReadFile("data.json")

	data := &Data{
		Data: []Game{},
    }

    json.Unmarshal(file, &data)

    newStruct := &Game{
		GameName: name,
		Path: txt,
    }

	fmt.Println(data)

	data.Data = append(data.Data, *newStruct)

	file, err := json.Marshal(data)
	if err != nil {
        fmt.Println(err)
    }
	err = ioutil.WriteFile("data.json", file, 0644)
	if err != nil {
        fmt.Println(err)
    }
}