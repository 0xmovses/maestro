package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/rvmelkonian/maestro/main/shared"
)

func (a *App) checkHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "OK!!")
}

func (a *App) createPitchMatrix(w http.ResponseWriter, r *http.Request) {
	var p shared.PitchRequest

	if err := scanJSON(r.Body, &p); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("pitches: %v \n", p)

	matrix := a.Composer.CreateToneRowMatrix(p.Pitches)
	respondWithJSON(w, http.StatusOK, matrix)
}

func (a *App) createVoicings(w http.ResponseWriter, r *http.Request) {
	var p shared.VoicingsRequest

	if err := scanJSON(r.Body, &p); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("pitches: %+v \n", p)

	voicings := a.Composer.CreateVoicings(p.Pitches, p.Amount)
	respondWithJSON(w, http.StatusOK, voicings)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func scanJSON(body io.ReadCloser, data interface{}) error {
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&data); err != nil {
		errMsg := "Invalid request payload."
		log.Error().Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}

	defer body.Close()

	return nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
