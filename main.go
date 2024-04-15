package main

import (
	"github.com/smarman85/gitTemplate/collaborator"
	"github.com/smarman85/gitTemplate/configureRepo"
	"github.com/smarman85/gitTemplate/newFromTemplate"
)

var (
	templateRepo = "name of the template repo"
	newRepo      = "name of the repo to create from template"
	org          = "org name these repos will be created under"
	adminGRP     = "group name for admin users"
	devGRP       = "group name for dev users"
)

func main() {
	r := newFromTemplate.Request{
		NewRepo:         newRepo,
		Template:        templateRepo,
		Owner:           org,
		Private:         true,
		IncludeBranches: true,
	}
	newFromTemplate.Build(r)

	c := configureRepo.TargetRepo{
		Repo:  newRepo,
		Owner: org,
	}
	configureRepo.Configure(c)

	// add admin collaborators
	adminInfo := collaborator.Request{
		TeamSlug:   adminGRP,
		Repository: newRepo,
		Org:        org,
		Permission: "admin",
	}
	collaborator.AddTeam(adminInfo)

	// adds dev collaborators
	devInfo := collaborator.Request{
		TeamSlug:   devGRP,
		Repository: newRepo,
		Org:        org,
		Permission: "maintain",
	}
	collaborator.AddTeam(devInfo)
}
