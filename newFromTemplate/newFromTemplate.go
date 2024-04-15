package newFromTemplate

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v55/github"
)

type Request struct {
	Template, NewRepo, Owner string
	client                   *github.Client
	ctx                      context.Context
	IncludeBranches, Private bool
}

type repo interface {
	verifyTemplate() error
	createFromTemplate() (*github.Repository, error)
}

func (r Request) verifyTemplate() error {

	repo, _, err := r.client.Repositories.Get(r.ctx, r.Owner, r.Template)
	if err != nil {
		return fmt.Errorf("error verifying template: %v", err)
	}

	fmt.Printf(
		"RepoName: %s\nTemplate: %t",
		*repo.Name,
		*repo.IsTemplate,
	)
	return nil
}

func (r Request) createFromTemplate() (*github.Repository, error) {
	repoRequest := github.TemplateRepoRequest{
		Name:               &r.NewRepo,
		Owner:              &r.Owner,
		Private:            &r.Private,
		IncludeAllBranches: &r.IncludeBranches,
	}
	repo, _, err := r.client.Repositories.CreateFromTemplate(r.ctx, r.Owner, r.Template, &repoRequest)
	if err != nil {
		return nil, fmt.Errorf("error creating repo from template: %v", err)
	}
	return repo, nil
}

func run(r repo) {
	err := r.verifyTemplate()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := r.createFromTemplate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
}

func Build(r Request) {
	r.client = github.NewClient(nil).WithAuthToken(os.Getenv("GIT_CLONE"))
	r.ctx = context.Background()
	run(r)
}
