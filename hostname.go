package main

// Author: Glen Newton
// BSD 3-Clause License

import (
	"github.com/hashicorp/golang-lru"
	"log"
	"net"
)

var hostnameCache *lru.Cache

func initCache() {
	var err error
	hostnameCache, err = lru.New(100)
	if err != nil {
		log.Println(err)
	}
}

// Returns dot at end of FQHN: http://www.dns-sd.org/trailingdotsindomainnames.html
//
func getRemoteHostname(ip string) string {
	emptyValue := EmptyValue

	if jsonOut {
		emptyValue = ""
	}

	if !lookupHostnames {
		return emptyValue
	}

	if hostnameCache.Contains(ip) {
		v, _ := hostnameCache.Get(ip)

		if hostname, ok := v.(string); ok {
			//fmt.Println("HIT")
			return hostname
		}
	}
	hostnames, err := net.LookupAddr(ip)
	if err != nil {
		return emptyValue
	} else {
		if len(hostnames) > 0 {
			hostnameCache.Add(ip, hostnames[0])
			return hostnames[0]
		}
	}
	return emptyValue

}
