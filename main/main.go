package main

import app "github.com/rvmelkonian/maestro/main/server"

func main() {
	a := app.App{}
	a.Initialize()
	a.Run()
}
