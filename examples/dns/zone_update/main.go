package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/hosting-de-labs/go-platform/client"
	"github.com/hosting-de-labs/go-platform/model"
)

var (
	flagBaseUrl     = flag.String("baseUrl", "https://secure.hosting.de/api/", "set api base url")
	flagZoneName    = flag.String("zoneName", "", "name of dns zone to update")
	flagRecordName  = flag.String("recordName", "", "name of record to create")
	flagRecordValue = flag.String("recordValue", "", "value of record to create")
)

func main() {
	flag.Parse()

	if len(*flagZoneName) == 0 || len(*flagRecordName) == 0 || len(*flagRecordValue) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	filter := &client.RequestFilter{
		Field: "ZoneNameUnicode",
		Value: *flagZoneName,
	}

	log.Printf("Searching for " + *flagZoneName)
	api := client.NewApiClient(*flagBaseUrl, os.Getenv("PLATFORM_TOKEN"), 250)
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
				Name:    *flagRecordName + "." + *flagZoneName,
				Type:    "A",
				Content: *flagRecordValue,
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
