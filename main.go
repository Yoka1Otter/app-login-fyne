package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "accounts.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var windowName string = "Авторизация"
	var windowNameNew string = "Панель"
	var labelText string = "Авторизируйтесь, чтобы продолжить."

	a := app.New()
	w := a.NewWindow(windowName)
	w.Resize(fyne.NewSize(300, 175))
	r, _ := LoadResourceFromPath("C:/Users/mydog/OneDrive/Рабочий стол/Golang/App_Login/myapp.png") // app icon
	w.SetIcon(r)
	label := widget.NewLabel(labelText)
	inputName := widget.NewEntry()
	inputPassword := widget.NewPasswordEntry()
	inputName.SetPlaceHolder("Введите логин...")
	inputPassword.SetPlaceHolder("Введите пароль...")
	buttonLogin := widget.NewButton("Авторизоваться", func() {
		username := inputName.Text
		password := inputPassword.Text

		db, err := sql.Open("sqlite3", "accounts.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		if inputName.Text == "" && inputPassword.Text == "" {
			fmt.Println("Введите данные")
		} else {
			query := "SELECT * FROM accountsdata WHERE name = ? AND pass = ?"
			var id int
			var name string
			var pass string
			err = db.QueryRow(query, username, password).Scan(&id, &name, &pass)
			if err != nil {
				fmt.Println("Неверное имя пользователя или пароль")
				return
			} else {
				w.Hide()
				w1 := a.NewWindow(windowNameNew)
				w1.Resize(fyne.NewSize(300, 175))
				w1.Show()
				fmt.Println("Все верно!")
			}
		}
	})

	w.SetContent(container.NewVBox(
		label,
		inputName,
		inputPassword,
		buttonLogin,
	))

	w.ShowAndRun()
}

type Resource interface {
	Name() string
	Content() []byte
}

type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

func (r *StaticResource) Name() string {
	return r.StaticName
}

func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}

func LoadResourceFromPath(path string) (Resource, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	name := filepath.Base(path)
	return NewStaticResource(name, bytes), nil
}
