package main

import (
	"github.com/smarman85/gitTemplate/newFromTemplate"
	// "github.com/smarman85/gitTemplate/configureRepo"
)

func main() {
	r := newFromTemplate.Request{
		NewRepo:  "newRepoFromTemplate",
		Tempalte: "template-repo",
		Owner:    "smarman85",
	}
	newFromTemplate.Build(r)
}
