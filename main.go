package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vault "example.com/vault-go/utils"

	"github.com/joho/godotenv"
)

var mountPath string = "secret"
var vaultPath string = "kubernetes-test"

func main() {

	err := godotenv.Load()
	vaultAddress := os.Getenv("vaultAddress")
	vaultToken := os.Getenv("vaultToken")
	ctx := context.Background()
	client, _ := vault.InitializeVaultCient(vaultAddress)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = client.SetToken(vaultToken)
	if err != nil {
		log.Println("Error initializing vault Client. Vault is not authenticated. Please login to vault using 'vault login -address", vaultAddress)
		panic(err)
	}
	fmt.Println("Vault client initialized")
	vault.WriteSecret(ctx, client, vaultPath)

}
