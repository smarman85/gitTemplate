package main

import (
	// "github.com/smarman85/gitTemplate/newFromTemplate"
	"github.com/smarman85/gitTemplate/configureRepo"
	"github.com/smarman85/gitTemplate/newFromTemplate"
)

func Configure() {
	r := newFromTemplate.Request{
		NewRepo:         "newRepoFromTemplate",
		Tempalte:        "template-repo",
		Owner:           "smarman85",
		Private:         false,
		IncludeBranches: false,
	}
	// newFromTemplate.Build(r)
	configureRepo.Configure(r)
}
