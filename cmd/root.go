package cmd

import (
	"os"

	"github.com/sebomancien/cloudflare/cmd/records"
	"github.com/sebomancien/cloudflare/internal/config"
	"github.com/spf13/cobra"
)

var Command cobra.Command

func init() {
	Command = cobra.Command{
		Use:   "cloudflare",
		Short: "Bundle of cloudflare tools",
		Long:  "Bundle of cloudflare tools",
	}

	Command.PersistentFlags().StringVarP(&config.Config.Email, "email", "e", os.Getenv("CF_EMAIL"), "cloudflare email")
	Command.PersistentFlags().StringVarP(&config.Config.ApiKey, "key", "k", os.Getenv("CF_API_KEY"), "cloudflare API key")

	Command.AddCommand(&records.Command)
}
