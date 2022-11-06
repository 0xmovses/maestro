package server

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/v1", a.checkHealth).Methods("GET")

	a.Router.HandleFunc("/api/v1/pitch/matrix", a.createPitchMatrix).Methods("POST")
	a.Router.HandleFunc("/api/v1/pitch/voicings", a.createVoicings).Methods("POST")
}
