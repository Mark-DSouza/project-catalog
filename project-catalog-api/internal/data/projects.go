package data

import (
	"time"

	"github.com/Mark-DSouza/project-catalog-api/internal/validator"
)

type Project struct {
	ID             int       `json:"id"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	RepositoryLink string    `json:"repository_link,omitzero"`
	DeploymentLink string    `json:"deployment_link,omitzero"`
	LifeStage      string    `json:"lifestage"`
	Viewability    string    `json:"viewability"`
}

func ValidProject(project *Project, v *validator.Validator) {
	v.Check(validator.NotBlank(project.Title), "title", "this field cannot be blank")
	v.Check(validator.MaxChar(project.Title, 100), "title", "this field must be at most 100")

	v.Check(validator.NotBlank(project.Description), "description", "this field cannot be blank")
	v.Check(validator.MaxChar(project.Description, 10), "description", "this field must be at least 10 characters long")
	v.Check(validator.MaxChar(project.Description, 1000), "description", "this field must be at most 1000 characters long")

	v.Check(validator.NotBlank(project.LifeStage), "lifestage", "this field cannot be blank")
	v.Check(validator.PermittedValue(project.LifeStage, PLANNING, IN_PROGRESS, COMPLETED, ARCHIVED), "lifestage", "this field must have one of the following values: PLANNING, IN-PROGRESS, COMPLETED, ARCHIVED")

	v.Check(validator.NotBlank(project.Viewability), "viewability", "this field cannot be blank")
	v.Check(validator.PermittedValue(project.Viewability, PUBLIC, PRIVATE), "viewability", "this field must have one of the following values: PRIVATE, PUBLIC")

}
