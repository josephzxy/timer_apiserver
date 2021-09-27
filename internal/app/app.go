package app

import "fmt"

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	fmt.Println("Running...")
}
