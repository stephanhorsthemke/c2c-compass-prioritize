package secrets

import (
	"context"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

const SECRETNAME string = "projects/826429547163/secrets/gdrive/versions/1"

var GDrive []byte

func init() {
	getGDriveSecret()
}

func getGDriveSecret() {
	// Create the client.
	ctx := context.Background()

	client, err := secretmanager.NewClient(ctx)
	defer client.Close()

	if err != nil {
		log.Fatal(err)
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: SECRETNAME,
	}

	resp, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get secret: %s", err)
	}

	GDrive = resp.Payload.Data
}
