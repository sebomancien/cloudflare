package main

import (
	"log"

	"github.com/sebomancien/cloudflare/cmd"
)

func main() {
	err := cmd.Command.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
