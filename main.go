package main

import (
	"log"
	"time"

	dns "github.com/Focinfi/go-dns-resolver"
	"github.com/joho/godotenv"
)

func main() {

	for {
		CheckDNS()
		time.Sleep(1 * time.Minute)
	}
}

func CheckDNS() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	domains := []string{"graph.salescloud.is"}
	types := []dns.QueryType{dns.TypeA, dns.TypeNS, dns.TypeMX, dns.TypeTXT}

	// Set timeout and retry times
	dns.Config.SetTimeout(uint(2))
	dns.Config.RetryTimes = uint(4)

	// Create and setup resolver with domains and types
	resolver := dns.NewResolver("1.1.1.1")
	resolver.Targets(domains...).Types(types...)
	// Lookup
	res := resolver.Lookup()

	//res.ResMap is a map[string]*ResultItem, key is the domain
	for target := range res.ResMap {
		log.Printf("%v: \n", target)
		for _, r := range res.ResMap[target] {
			if r.Type == "A" {
				SendMessageToDiscord(r.Record + " " + r.Type + " " + r.Content + " " + time.Now().Format(time.RFC3339))
			}
		}
	}
}
