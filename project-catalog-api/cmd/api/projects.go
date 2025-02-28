package main

import (
	"net/http"
	"time"

	"github.com/Mark-DSouza/project-catalog-api/internal/data"
)

func (app *application) showProjectHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	project := data.Project{
		ID:             int(id),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Title:          "Project Catalog",
		Description:    "A project catalog to keep track of your personal projects in Go and React",
		RepositoryLink: "www.github.com/Mark-DSouza/project-catalog",
		DeploymentLink: "",
		LifeStage:      data.IN_PROGRESS,
		Viewability:    data.PUBLIC,
	}

	err = app.writeJSON(w, http.StatusOK, project, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
