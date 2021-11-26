package main

import (
	"fmt"
	"log"
	"time"

	dns "github.com/Focinfi/go-dns-resolver"
	"github.com/joho/godotenv"
	"golang.org/x/net/route"
)

func main() {

	for {
		CheckDNS()
		time.Sleep(1 * time.Minute)
	}
}

var defaultRoute = [4]byte{0, 0, 0, 0}

func GetGateway() {
	rib, _ := route.FetchRIB(0, route.RIBTypeRoute, 0)
	messages, err := route.ParseRIB(route.RIBTypeRoute, rib)

	if err != nil {
		return
	}

	for _, message := range messages {
		route_message := message.(*route.RouteMessage)
		addresses := route_message.Addrs

		var destination, gateway *route.Inet4Addr
		ok := false

		if destination, ok = addresses[0].(*route.Inet4Addr); !ok {
			continue
		}

		if gateway, ok = addresses[1].(*route.Inet4Addr); !ok {
			continue
		}

		if destination == nil || gateway == nil {
			continue
		}

		if destination.IP == defaultRoute {
			fmt.Println(gateway.IP)
		}
	}
}

func CheckDNS() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	domains := []string{"salescloud.is"}
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
