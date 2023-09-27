package main

import (
	"github.com/smarman85/gitTemplate/configureRepo"
	_ "github.com/smarman85/gitTemplate/newFromTemplate"
)

func main() {
	/*
		r := newFromTemplate.Request{
			NewRepo:         "newRepoFromTemplate",
			Tempalte:        "template-repo",
			Owner:           "smarman85",
			Private:         false,
			IncludeBranches: false,
		}
		newFromTemplate.Build(r)
	*/

	c := configureRepo.TargetRepo{
		Repo:  "newRepoFromTemplate",
		Owner: "smarman85",
	}
	configureRepo.Configure(c)
}
