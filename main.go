package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/compute/metadata"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/gcp-secret-manager-go/impersonate"
	"github.com/spf13/viper"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var (
	client *secretmanager.Client
	err    error
)

func main() {
	// Create context.
	ctx := context.Background()

	// Create client.
	client, err = secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// If local, reinitialize client with impersonation details.
	if !metadata.OnGCE() {
		client, err = impersonate.Impersonate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Request payload.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/pantheon-lighthouse-poc/secrets/test-sachin-vrt-660/versions/latest",
	}

	// Send request.
	resp, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	// Viper setting.
	viper.SetConfigType("json")
	// Reading config.
	if err := viper.ReadConfig(bytes.NewBuffer(resp.Payload.Data)); err != nil {
		log.Fatal(err)
	}

	// Print.
	prettyConfig, err := json.MarshalIndent(viper.AllSettings(), "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("viper-config: %s", string(prettyConfig))
}
