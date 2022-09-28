package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"

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

	path := "./games"
	tab := []string{}
	gn := Search_game(path,tab)
	Display_game(w,s,gn)
	w.Resize(fyne.NewSize(800,400))
	// a.Settings().SetTheme(&myTheme{})
	w.ShowAndRun()
}


func Search_game(path string, tab []string) []string{
	items, _ := ioutil.ReadDir(path)
    for _, item := range items {
		//fmt.Println(item.Name())
		match, _ := regexp.MatchString(".exe$", item.Name())
		if match {
			fmt.Println(item.Name())
			tab = append(tab,item.Name())
		} else if item.IsDir(){
			tab = In_folder("./games/"+item.Name(),tab)
		}
	}
	return tab
}

func In_folder(path string, tab []string) []string {
	items, _ := ioutil.ReadDir(path)
    for _, item := range items {
		match, _ := regexp.MatchString(".exe$", item.Name())
		if match {
			tab = append(tab,path + "/" + item.Name())
			return tab
		} else if item.IsDir(){
			tab = In_folder(path + "/" +item.Name(),tab)
		}
	}
	return tab
}

func Display_game(w fyne.Window, s *fyne.Container, game_path []string){

		var s2 *fyne.Container
		var box *container.Scroll
		s2 = container.New(layout.NewVBoxLayout())

		//display games layout
		for _, i := range game_path {
			fmt.Println(i)
			game_name := strings.Split(i, "/")
			name := game_name[len(game_name)-1]
			var s3 *fyne.Container
			s3 = container.New(layout.NewGridLayout(3))
			label := widget.NewLabel(name)
			x := widget.NewButton("lanch", func() {lauch_game(i)})
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

func lauch_game(path string){
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