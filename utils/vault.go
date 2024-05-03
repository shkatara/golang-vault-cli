package utils

import (
	"context"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

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

func WriteSecret(ctx context.Context, client *vault.Client, vaultPath string, certsDir string) error {
	_, err := client.Secrets.KvV2Write(ctx, vaultPath, schema.KvV2WriteRequest{
		Data: map[string]any{
			"password1": "abcasdasda123",
			"password2": "correct horse battery staple",
		}},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		return err
	}
	return nil
}
