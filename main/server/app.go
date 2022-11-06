package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rvmelkonian/maestro/main/maestro"
	"github.com/rvmelkonian/maestro/main/middleware"
	"github.com/rvmelkonian/maestro/main/shared"
)

type App struct {
	Router *mux.Router
	DB     *shared.Database

	Env      string
	Config   shared.Config
	Composer maestro.PitchMake
}

func (a *App) Run() {
	addr := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Info().Msgf("Starting server on %s ...", addr)
	log.Fatal().Err(http.ListenAndServe(addr, a.Router)).Msgf("Server at %s crashed!", addr)
}

func (a *App) Initialize() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	////////////
	// Clients
	////////////

	// Postgres
	//dbname := os.Getenv("DB_NAME")

	// Postgres
	// a.ConnectDB(
	// 	os.Getenv("DB_USERNAME"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	dbname,
	// )

	// Router
	a.Router = mux.NewRouter()
	a.initializeRoutes()

	// Middlewares
	a.Router.Use(mux.CORSMethodMiddleware(a.Router))
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.UseCors(a.Config))

	// Composer
	a.Composer = maestro.PitchMake{}
}

func (a *App) ConnectDB(username, password, host, port, dbname string) {
	var database shared.Database
	var err error

	database.Context = context.Background()
	database.Name = dbname

	connectionString :=
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)

	pconf, confErr := pgxpool.ParseConfig(connectionString)
	if confErr != nil {
		log.Fatal().Err(err).Msg("Unable to parse database config url")
	}

	if os.Getenv("APP_ENV") == "TEST" {
		log.Info().Msg("Setting MIN/MAX connections to 1")
		pconf.MinConns = 1
		pconf.MaxConns = 1
	}

	database.Conn, err = pgxpool.ConnectConfig(database.Context, pconf)

	database.Env = &a.Env
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating Postsgres conn pool")
	} else {
		a.DB = &database
		log.Info().Msgf("Successfully created Postgres conn pool")
	}
}
