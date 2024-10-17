package ls

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sebomancien/cloudflare/internal/config"
	"github.com/sebomancien/cloudflare/internal/dns/cloudflare"
	"github.com/spf13/cobra"
)

var Command cobra.Command

func init() {
	Command = cobra.Command{
		Use:   "ls",
		Short: "List all records",
		Long:  "List all records in the cloudflare zone",
		Run:   command,
	}
}

func command(cmd *cobra.Command, args []string) {
	records, err := cloudflare.GetRecords(&config.Config)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range records {
		json, err := json.Marshal(r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(json))
	}
}
