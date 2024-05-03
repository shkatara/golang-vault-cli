package main

import (
	"context"
	"flag"
	"log"
	"os"

	utils "example.com/vault-go/utils"
	vault_utils "example.com/vault-go/vault"

	"github.com/hashicorp/vault-client-go"
	"github.com/joho/godotenv"
)

var (
	vaultPath string = "kubernetes-test"
	domain    string
)

func init_info() (context.Context, *vault.Client) {
	err := godotenv.Load()
	utils.Checkerr(err)

	vaultAddress := os.Getenv("vaultAddress")
	vaultToken := os.Getenv("vaultToken")
	ctx := context.Background()

	client, _ := vault_utils.InitializeVaultCient(vaultAddress)
	err = client.SetToken(vaultToken)
	if err != nil {
		log.Println("Error initializing vault Client. Vault is not authenticated. Please login to vault using 'vault login -address", vaultAddress)
		panic(err)
	}
	return ctx, client
}
func main() {
	flag.StringVar(&domain, "domain", domain, "Directory to read metrics from")
	flag.Parse()
	ctx, client := init_info()
	vault_utils.WriteSecret(ctx, client, vaultPath, domain)

}
