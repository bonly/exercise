package main

import (
	"fmt"
	"github.com/abh/geoip"
)

func main() {
	gi, err := geoip.Open("GeoIP.dat");
	if err != nil {
		fmt.Printf("Could not open GeoIP database: %s\n", err);
	}
	if gi != nil {
		test4(*gi, "207.171.7.51");
		test4(*gi, "10.47.173.157");
		test4(*gi, "8.8.8.8");
	}
	rc, _:= gi.GetName("202.13.34.13");
	fmt.Println(rc);
}

func test4(g geoip.GeoIP, ip string) {
	test(func(s string) (string, int) { return g.GetCountry(s) }, ip);
}

func test(f func(string) (string, int), ip string) {
	country, netmask := f(ip);
	fmt.Printf("ip: %s is [%s] (netmask %d)\n", ip, country, netmask);

}
