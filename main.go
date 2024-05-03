package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/joho/godotenv"
)

var mountPath string = "secret"
var vaultPath string = "kubernetes-test"

func initializeVaultCient(address string) (*vault.Client, error) {
	vaultClient, err := vault.New(
		vault.WithAddress(address),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		return nil, err
	}
	return vaultClient, nil
}

func writeSecret(ctx context.Context, client *vault.Client) {
	_, err := client.Secrets.KvV2Write(ctx, vaultPath, schema.KvV2WriteRequest{
		Data: map[string]any{
			"password1": "abc123",
			"password2": "correct horse battery staple",
		}},
		vault.WithMountPath(mountPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written successfully")
}

func main() {
	vaultAddress := os.Getenv("vaultAddress")
	vaultToken := os.Getenv("vaultToken")
	ctx := context.Background()
	client, _ := initializeVaultCient(vaultAddress)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client.SetToken(vaultToken)
	if err != nil {
		log.Println("Error initializing vault Client. Vault is not authenticated. Please login to vault using 'vault login -address", vaultAddress)
		panic(err)
	}
	fmt.Println("Vault client initialized")
	writeSecret(ctx, client)

}
