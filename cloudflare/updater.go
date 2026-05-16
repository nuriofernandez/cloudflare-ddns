package cloudflare

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/dns"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

var apiToken = os.Getenv("CLOUDFLARE_API_TOKEN")
var dnsRecordId = os.Getenv("CLOUDFLARE_DNS_RECORD_ID")
var zoneId = os.Getenv("CLOUDFLARE_ZONE_ID")
var recordName = os.Getenv("CLOUDFLARE_RECORD_NAME")

func Update(address string) error {
	client := cloudflare.NewClient(
		option.WithAPIToken(apiToken),
	)
	recordResponse, err := client.DNS.Records.Update(
		context.TODO(),
		dnsRecordId,
		dns.RecordUpdateParams{
			ZoneID: cloudflare.F(zoneId),
			Body: dns.ARecordParam{
				Name:    cloudflare.F(recordName),
				Proxied: cloudflare.F(false),
				Content: cloudflare.F(address),
				TTL:     cloudflare.F(dns.TTL1),
				Type:    cloudflare.F(dns.ARecordTypeA),
			},
		},
	)

	if err != nil {
		fmt.Printf("[Cloudflare] Error updating record: %s\n", recordName)
		fmt.Printf("%+v\n", recordResponse)
		return err
	}

	fmt.Printf("[Cloudflare] Updated record: %s\n", recordName)
	return nil
}
