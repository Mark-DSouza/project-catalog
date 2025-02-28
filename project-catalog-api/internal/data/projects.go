package data

import "time"

type Project struct {
	ID             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Title          string
	Description    string
	RepositoryLink string
	DeploymentLink string
	LifeStage      string
	Viewability    string
}
