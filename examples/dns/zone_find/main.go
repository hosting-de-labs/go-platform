package main

import (
	"log"
	"os"

	"github.com/hosting-de-labs/go-platform/client"
)

func main() {
	filter := &client.RequestFilter{
		Field: "zoneName",
		Value: "aac1.dc.zone",
	}

	c := client.NewApiClient("https://secure.hosting.de/api/", os.Getenv("PLATFORM_TOKEN"), 250)
	res, err := c.Dns.ZonesFind(filter)
	if err != nil {
		panic(err)
	}

	for _, r := range res {
		log.Printf("%s", r.ZoneConfig.ID)
	}
}
