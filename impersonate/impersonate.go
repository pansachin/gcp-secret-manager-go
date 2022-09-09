package impersonate

import (
	"context"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

func Impersonate(ctx context.Context) (*secretmanager.Client, error) {
	// Get token source

	ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: "test-vrt-660@pantheon-lighthouse-poc.iam.gserviceaccount.com",
		Scopes:          []string{"https://www.googleapis.com/auth/cloud-platform"},
		// Delegates: []string{},
	})
	if err != nil {
		return nil, err
	}
	client, err := secretmanager.NewClient(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	return client, nil
}
