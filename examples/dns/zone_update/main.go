package main

import (
	"log"
	"os"
	"time"

	"github.com/hosting-de-labs/go-platform/client"
	"github.com/hosting-de-labs/go-platform/model"
)

func main() {
	api := client.NewApiClient("https://secure.hosting.de/api/", os.Getenv("PLATFORM_TOKEN"), 250)

	filter := &client.RequestFilter{
		Field: "ZoneNameUnicode",
		Value: "*keenlogics.dev*",
	}

	log.Printf("Searching for keenlogics.dev")

	zcs, err := api.Dns.ZoneConfigsFind(filter)
	if err != nil {
		panic(err)
	}

	if len(zcs) > 1 {
		panic("More than 1 result received")
	}

	ttl := 1 * time.Hour
	req := client.ZoneUpdateRequest{
		ZoneConfig: zcs[0],
		RecordsToAdd: []model.RecordObject{
			{
				Name:    "hsokolowski-test.keenlogics.dev",
				Type:    "A",
				Content: "127.0.0.1",
				TTL:     int(ttl.Seconds()),
			},
		},
	}

	log.Printf("Found zone with id %s, adding record", zcs[0].AccountID)

	res, err := api.Dns.ZoneUpdate(req)
	if err != nil {
		panic(err)
	}

	if len(res.Errors) > 0 {
		log.Printf("Errors occured:")
		for _, e := range res.Errors {
			log.Printf("%s", e)
		}

		log.Fatal("Exiting...")
	}

	log.Printf("New last change date of zone: %s", res.Response.ZoneConfig.LastChangeDate)
}
