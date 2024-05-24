package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	utils "example.com/vault-go/utils"
	vault_utils "example.com/vault-go/vault"

	"github.com/hashicorp/vault-client-go"
	"github.com/joho/godotenv"
)

var (
	vaultPath  string = "ssl-certs"
	domain     string
	vaultMount string = "secret"
)

func init_info() (context.Context, *vault.Client) {
	err := godotenv.Load()
	utils.Checkerr(err)
	vaultAddress := os.Getenv("vaultAddress")
	vaultTokenPath := fmt.Sprintf("/Users/%s/.vault-token", os.Getenv("USER"))
	vaultTokenByte, _ := os.ReadFile(vaultTokenPath)
	vaultToken := string(vaultTokenByte)
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
	vault_utils.WriteSecret(ctx, client, vaultPath, vaultMount, domain)
}
