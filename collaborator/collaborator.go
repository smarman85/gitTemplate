package collaborator

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

type collaborator interface {
	addCollaborator()
}

type team interface {
	addTeam()
}

type Request struct {
	UserName, Repository, Org, Permission, TeamSlug string
	client                                          *github.Client
	ctx                                             context.Context
}

func setClient() *github.Client {
	token := os.Getenv("GIT_CLONE")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func (r Request) addCollaborator() {
	log.Printf(
		"Adding: %s to %s:%s with %s permissions",
		r.UserName,
		r.Org,
		r.Repository,
		r.Permission,
	)
	opts := &github.RepositoryAddCollaboratorOptions{
		Permission: r.Permission,
	}
	_, _, err := r.client.Repositories.AddCollaborator(
		r.ctx,
		r.Org,
		r.Repository,
		r.UserName,
		opts,
	)
	if err != nil {
		fmt.Println(err)
	}
}

func (r Request) addTeam() {
	log.Printf(
		"Adding: %s to %s:%s with %s permissions",
		r.TeamSlug,
		r.Org,
		r.Repository,
		r.Permission,
	)

	opts := &github.TeamAddTeamRepoOptions{
		Permission: r.Permission,
	}

	_, err := r.client.Teams.AddTeamRepoBySlug(
		r.ctx,
		r.Org,
		r.TeamSlug,
		r.Org,
		r.Repository,
		opts,
	)

	if err != nil {
		log.Printf("error inviting team: %v", err)
	}
}

func runUser(c collaborator) {
	c.addCollaborator()
}

func runTeam(t team) {
	t.addTeam()
}

/*
func validPermissions(r Request) bool {
	// Possible values are:
	//     pull - team members can pull, but not push to or administer this repository
	//     push - team members can pull and push, but not administer this repository
	//     admin - team members can pull, push and administer this repository
	//     maintain - team members can manage the repository without access to sensitive or destructive actions.
	//     triage - team members can proactively manage issues and pull requests without write access.
}
*/

func AddUser(r Request) {
	r.ctx = context.Background()
	r.client = setClient()
	runUser(r)
	/*
		if validPermissions(r) && r.UserName != "" {
			r.ctx = context.Background()
			r.client = setClient()
			runUser(r)
		} else {
			log.Fatalf("Invalid permission: %s for user: %s", r.Permission, r.UserName)
		}
	*/
}

func AddTeam(r Request) {
	r.ctx = context.Background()
	r.client = setClient()
	runTeam(r)
	/*
		if validPermissions(r) && r.TeamSlug != "" {
			r.ctx = context.Background()
			r.client = setClient()
			runTeam(r)
		} else {
			log.Fatalf("Invalid permission: %s for user: %s", r.Permission, r.UserName)
		}
	*/
}
