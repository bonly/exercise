package main

import (
	"fmt"
	"github.com/abh/geoip"
)

func main() {

	file6 := "./GeoIPv6.dat"

	gi6, err := geoip.Open(file6)
	if err != nil {
		fmt.Printf("Could not open GeoIPv6 database: %s\n", err)
	}

	gi, err := geoip.Open()
	if err != nil {
		fmt.Printf("Could not open GeoIP database: %s\n", err)
	}

	giasn, err := geoip.Open("./GeoIPASNum.dat")
	if err != nil {
		fmt.Printf("Could not open GeoIPASN database: %s\n", err)
	}

	giasn6, err := geoip.Open("./GeoIPASNumv6.dat")
	if err != nil {
		fmt.Printf("Could not open GeoIPASN database: %s\n", err)
	}

    gcity, err := geoip.Open("./GeoLiteCity.dat");
	if err != nil {
		fmt.Printf("Could not open GeoIPASN database: %s\n", err)
	}

	if giasn != nil {
		ip := "207.171.7.51"
		asn, netmask := giasn.GetName(ip)
		fmt.Printf("%s: %s (netmask /%d)\n", ip, asn, netmask)

	}

	if gcity != nil{
		ip := "10.47.173.157"
		country, _ := gcity.GetRegion(ip);
		fmt.Printf("%s: %s\n", ip, country);
	}

	if gi != nil {
		test4(*gi, "207.171.7.51")
		test4(*gi, "127.0.0.1")
	}
	if gi6 != nil {
		ip := "2607:f238:2::5"
		country, netmask := gi6.GetCountry_v6(ip)
		var asn string
		var asn_netmask int
		if giasn6 != nil {
			asn, asn_netmask = giasn6.GetNameV6(ip)
		}
		fmt.Printf("%s: %s/%d %s/%d\n", ip, country, netmask, asn, asn_netmask)

	}

}

func test4(g geoip.GeoIP, ip string) {
	test(func(s string) (string, int) { return g.GetCountry(s) }, ip)
}

func test(f func(string) (string, int), ip string) {
	country, netmask := f(ip)
	fmt.Printf("ip: %s is [%s] (netmask %d)\n", ip, country, netmask)

}

/*
  http://geolite.maxmind.com/download/geoip/database/GeoLiteCountry/GeoIP.dat.gz
  http://geolite.maxmind.com/download/geoip/database/GeoIPv6.dat.gz
  http://geolite.maxmind.com/download/geoip/database/GeoLiteCity.dat.gz
  http://geolite.maxmind.com/download/geoip/database/GeoLiteCityv6-beta/GeoLiteCityv6.dat.gz
  http://download.maxmind.com/download/geoip/database/asnum/GeoIPASNum.dat.gz
  http://download.maxmind.com/download/geoip/database/asnum/GeoIPASNumv6.dat.gz
*/