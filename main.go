package main

import (
	_ "github.com/joho/godotenv/autoload" // Important to keep as first import!

	"fmt"
	"log"

	"ddns/cloudflare"
	"ddns/ipinfo"
	"ddns/scheduler"
)

func main() {
	scheduler.Schedule(func() {
		fmt.Println("[DDNS] Attempting to update DDNS...")
		ip, err := ipinfo.FetchIp()
		if err != nil {
			fmt.Println("[DDNS] Unable to fetch current IP address.")
			log.Fatal(err)
		}

		err = cloudflare.Update(ip)
		if err != nil {
			fmt.Println("[DDNS] Unable to update Cloudflare DNS record.")
			log.Fatal(err)
		}

		fmt.Println("[DDNS] Successfully updated DDNS record!")
	})
}
