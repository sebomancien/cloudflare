package update

import (
	"log"

	"github.com/sebomancien/cloudflare/internal/config"
	"github.com/sebomancien/cloudflare/internal/dns/cloudflare"
	"github.com/sebomancien/cloudflare/internal/utils"
	"github.com/spf13/cobra"
)

var Command cobra.Command

func init() {
	Command = cobra.Command{
		Use:   "update",
		Short: "Update all records",
		Long:  "Updates all records to the current public address",
		Run:   command,
	}
}

func command(cmd *cobra.Command, args []string) {
	ipAddress, err := utils.GetPublicIp()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current Public IP:", ipAddress)

	records, err := cloudflare.GetRecords(&config.Config)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range records {
		if r.Type != "A" {
			continue
		}

		if r.Content != ipAddress {
			log.Println("Updating record id:", r.Name)
			r.Content = ipAddress
			err = cloudflare.PutRecord(r, &config.Config)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("All records are up to date")
}
