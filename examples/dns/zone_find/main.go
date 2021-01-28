package main

import (
	"log"

	"github.com/hosting-de-labs/go-platform/client"
)

func main() {
	filter := &client.RequestFilter{
		Field: "zoneName",
		Value: "aac1.dc.zone",
	}

	c := client.NewApiClient("https://secure.hosting.de/api/", "<API_KEY>", 250)
	res, err := c.ZonesFind(filter)
	if err != nil {
		panic(err)
	}

	for _, r := range res {
		log.Printf("%s", r.ZoneConfig.ID)
	}
}
