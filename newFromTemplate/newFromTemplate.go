package newFromTemplate

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v55/github"
)

type Request struct {
	Tempalte, NewRepo, Owner string
	client                   *github.Client
	ctx                      context.Context
}

type repo interface {
	getRepoInfo()
}

func (r Request) getRepoInfo() {

	repo, _, err := r.client.Repositories.Get(r.ctx, r.Owner, r.Tempalte)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(
		"RepoName: %s\nTemplate: %t",
		*repo.Name,
		*repo.IsTemplate,
	)
}

func run(r repo) {
	r.getRepoInfo()
}

func Build(r Request) {
	r.client = github.NewClient(nil).WithAuthToken(os.Getenv("GIT_CLONE"))
	r.ctx = context.Background()
	run(r)
}
