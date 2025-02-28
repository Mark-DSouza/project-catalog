package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mark-DSouza/project-catalog-api/internal/data"
	"github.com/Mark-DSouza/project-catalog-api/internal/validator"
)

func (app *application) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		RepositoryLink string `json:"repository_link"`
		DeploymentLink string `json:"deployment_link"`
		LifeStage      string `json:"lifestage"`
		Viewability    string `json:"viewability"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	project := &data.Project{
		Title:          input.Title,
		Description:    input.Description,
		RepositoryLink: input.RepositoryLink,
		DeploymentLink: input.DeploymentLink,
		LifeStage:      input.LifeStage,
		Viewability:    input.Viewability,
	}

	v := validator.New()
	data.ValidProject(project, v)

	if !v.Valid() {
		app.validationErrorResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", project)
}

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
