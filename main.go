package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct {}


func Display_laucher(){
	// var _ fyne.Theme = (*myTheme)(nil)
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
	// a.Settings().SetTheme(&myTheme{})
	w.ShowAndRun()
}


func Display_game(w fyne.Window, s *fyne.Container){

		var s2 *fyne.Container
		var box *container.Scroll
		s2 = container.New(layout.NewVBoxLayout())

		    //read data.json
			f, _ := ioutil.ReadFile("data.json")
			datas := Data{}
			_ = json.Unmarshal([]byte(f), &datas)

		//display games layout
		for i := 0; i < len(datas.Data); i++ {
			fmt.Println(datas.Data[i])
			var s3 *fyne.Container
			s3 = container.New(layout.NewGridLayout(3))
			label := widget.NewLabel(datas.Data[i].GameName)
			fn := func(i int) func() {
                return func() {
                    launch_game(datas.Data[i].Path)
                }
            }(i)
			x := widget.NewButton("launch", fn)
			s3.Add(label)
			s3.Add(x)

			s2.Add(s3)
		}

		box = container.NewVScroll(s2)
		box.SetMinSize(fyne.NewSize(100,400))

		w.SetContent(container.NewVBox(
			s,
			box,
		),)
}

func launch_game(path string){
    cmnd := exec.Command(path, "arg")
    cmnd.Start()
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func search(){

}

func main() {
	Display_laucher()
}