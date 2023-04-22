package main

import (
	"net/http"
)

type Env struct {
	Status     string     `json:"status"`
	SystemInfo SystemInfo `json:"system_info"`
}

type SystemInfo struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

//	@Summary		Healthcheck
//	@Description	HealthCheck of server
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Env
//	@Failure		500	{object}	Error
//	@Router			/healthcheck [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
