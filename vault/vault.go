package vault

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/joho/godotenv"
)

var certsDir string = "certs"

func InitializeVaultCient(address string) (*vault.Client, error) {
	vaultClient, err := vault.New(
		vault.WithAddress(address),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		return nil, err
	}
	return vaultClient, nil
}

func WriteSecret(ctx context.Context, client *vault.Client, vaultPath string, vaultMount string, domain string) {
	godotenv.Load()
	serverCrt, _ := os.ReadFile(certsDir + "/" + domain + "_server.crt")
	intermediate, _ := os.ReadFile(certsDir + "/" + domain + "_intermediate.crt")
	root, _ := os.ReadFile(certsDir + "/" + domain + "_root.crt")
	key, _ := os.ReadFile(certsDir + "/" + domain + "_server.key")

	_, err := client.Secrets.KvV2Write(ctx, vaultPath, schema.KvV2WriteRequest{
		Data: map[string]any{
			domain + "_server.crt":       string(serverCrt),
			domain + "_intermediate.crt": string(intermediate),
			domain + "_root.crt":         string(root),
			domain + "_server.key":       string(key),
		}},
		vault.WithMountPath(vaultMount),
	)
	if err != nil {
		log.Println("Error writing secret to vault")
		log.Fatal(err)
	}
	log.Println("Secret for domain", domain, "written successfully to", os.Getenv("vaultAddress"))
}
