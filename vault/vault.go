package vault

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
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

func WriteSecret(ctx context.Context, client *vault.Client, vaultPath string, domain string) {

	serverCrt, _ := os.ReadFile(certsDir + "/" + domain + ".crt")
	intermediate, _ := os.ReadFile(certsDir + "/" + domain + ".intermediate")
	key, _ := os.ReadFile(certsDir + "/" + domain + ".key")
	log.Print(string(serverCrt))

	_, err := client.Secrets.KvV2Write(ctx, vaultPath, schema.KvV2WriteRequest{
		Data: map[string]any{
			domain + ".crt":          string(serverCrt),
			domain + ".intermediate": string(intermediate),
			domain + ".key":          string(key),
		}},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written successfully")
}
