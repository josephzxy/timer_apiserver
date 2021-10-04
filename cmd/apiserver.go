package main

import "github.com/josephzxy/timer_apiserver/internal/app"

func main() {
	app.NewApp("apiserver").Run()
}
