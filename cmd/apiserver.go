// apiserver is the API server that provides both external-facing RESTful API
// and internal-facing gRPC API for managing RESTful resource "Timer".
package main

import "github.com/josephzxy/timer_apiserver/internal/app"

func main() {
	app.New("apiserver").Run()
}
