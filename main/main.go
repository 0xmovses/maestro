package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	app "github.com/rvmelkonian/maestro/main/server"
)

func main() {
	var CWD string
	var err error

	if CWD = os.Getenv("CWD"); len(CWD) == 0 {
		CWD = "./"
		os.Setenv("CWD", CWD)
	}

	if UseDotEnv := os.Getenv("ENV_NAME"); len(UseDotEnv) == 0 {
		err = godotenv.Load(CWD + ".env")
	}
	if err != nil {
		log.Fatalf("Error loading .env file!!!\n")
	}
	a := app.App{}
	a.Initialize()
	a.Run()
}
