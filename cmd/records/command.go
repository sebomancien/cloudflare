package records

import (
	"os"

	"github.com/sebomancien/cloudflare/cmd/records/ls"
	"github.com/sebomancien/cloudflare/cmd/records/update"
	"github.com/sebomancien/cloudflare/internal/config"
	"github.com/spf13/cobra"
)

var Command cobra.Command

func init() {
	Command = cobra.Command{
		Use:   "records",
		Short: "Commands related to DNS records",
		Long:  "Commands related to DNS records",
	}

	Command.PersistentFlags().StringVarP(&config.Config.ZoneId, "zone", "z", os.Getenv("CF_ZONE_ID"), "cloudflare zone id")

	Command.AddCommand(&ls.Command)
	Command.AddCommand(&update.Command)
}
