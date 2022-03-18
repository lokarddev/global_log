package main

import "github.com/lokarddev/global_log/cmd/app"

func main() {
	application := app.NewApplication()
	application.InitApp()
	application.Run()
}
