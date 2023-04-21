package ip2proxy_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/atreides-intelligence/ip2proxy-go"
)

func TestIP2L(m *testing.T) {
	ip := "8.8.8.8"

	proxy, _ := ioutil.ReadFile("./IP2PROXY-LITE-PX7.BIN")

	// Open the GeoLite2 ASN database
	proxydb, err := ip2proxy.OpenDBWithBytes(proxy)
	if err != nil {
		panic(err)
	}

	proxyInfo, err := proxydb.GetAll(ip)
	if err != nil {
		panic(err)
	}

	// Print results
	fmt.Println("IP:", ip)
	x2, _ := json.Marshal(proxyInfo)
	fmt.Println(string(x2))
}
