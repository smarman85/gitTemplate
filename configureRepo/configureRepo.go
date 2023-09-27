package configureRepo

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v55/github"
)

type repo interface {
	setBranchProtection()
	// createPrivatePullSecrets()
}

type TargetRepo struct {
	Repo, Owner string
	client      *github.Client
	ctx         context.Context
}

func (t TargetRepo) setBranchProtection() {
	rule := github.ProtectionRequest{
		// RequiredStatusChecks: "",
		RequiredPullRequestReviews: &github.PullRequestReviewsEnforcementRequest{
			RequireCodeOwnerReviews: true,
		},
		EnforceAdmins: false,
		// Restrictions: "",

	}
	_, r, err := t.client.Repositories.UpdateBranchProtection(t.ctx, t.Owner, t.Repo, "main", &rule)
	if err != nil {
		fmt.Errorf("error setting branch protection rule: %v", err)
	}

	// fmt.Println(p)
	fmt.Println(r.Response.StatusCode)
}

/*
func (t TargetRepo) createPrivatePullSecrets() {
	secret := github.EncryptedSecret{}
	r, err := t.client.Actions.CreateOrUpdateRepoSecret(t.ctx, t.Owner, t.Repo, secret)
	if err != nil {
		fmt.Errorf("error setting repo secret: %v", err)
	}
	fmt.Println(r.Request.Response.StatusCode)

}
*/

func run(r repo) {
	r.setBranchProtection()
}

func Configure(t TargetRepo) {
	t.client = github.NewClient(nil).WithAuthToken(os.Getenv("GIT_CLONE"))
	t.ctx = context.Background()
	run(t)
}
