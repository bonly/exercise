package main

import (
    "fmt"
    "github.com/oschwald/geoip2-golang"
    "net"
)

func main() {
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "183.61.112.2";
        ip := net.ParseIP(str)
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        fmt.Printf("ip: %s\n", str);
        fmt.Printf("城市: %v\n", record.City.Names["zh-CN"]);
        fmt.Printf("国家: %v\n", record.Country.Names["zh-CN"]);
        fmt.Printf("省份: %v\n", record.Subdivisions[0].Names["zh-CN"]);
        fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
        fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
        fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
        fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
        fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
        fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
        fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
    }
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "172.33.44.11";
        ip := net.ParseIP(str);
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        if len(record.Subdivisions)>0 {
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v \t 省份: %s\n", 
                str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"],record.Subdivisions[0].Names["zh-CN"]);
        }else{
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v\n", str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"]);
        }
    }
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "113.108.208.85";
        ip := net.ParseIP(str);
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        if len(record.Subdivisions)>0 {
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v \t 省份: %s\n", 
                str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"],record.Subdivisions[0].Names["zh-CN"]);
        }else{
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v\n", str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"]);
        }
    }
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "8.8.8.8";
        ip := net.ParseIP(str);
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        if len(record.Subdivisions)>0 {
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v \t 省份: %s\n", 
                str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"],record.Subdivisions[0].Names["zh-CN"]);
        }else{
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v\n", str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"]);
        }
    }    
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "10.0.0.1";
        ip := net.ParseIP(str);
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        if len(record.Subdivisions)>0 {
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v \t 省份: %s\n", 
                str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"],record.Subdivisions[0].Names["zh-CN"]);
        }else{
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v\n", str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"]);
        }
    }    
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "9.176.122.157";
        ip := net.ParseIP(str);
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        if len(record.Subdivisions)>0 {
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v \t 省份: %s\n", 
                str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"],record.Subdivisions[0].Names["zh-CN"]);
        }else{
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v\n", str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"]);
        }
    }        
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "4.69.149.18";
        ip := net.ParseIP(str);
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        if len(record.Subdivisions)>0 {
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v \t 省份: %s\n", 
                str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"],record.Subdivisions[0].Names["zh-CN"]);
        }else{
            fmt.Printf("ip: %s \t 国家: %v \n", str, record.Country.Names["zh-CN"]);
        }
    } 
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "192.29.2.199";
        ip := net.ParseIP(str);
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        if len(record.Subdivisions)>0 {
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v \t 省份: %s\n", 
                str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"],record.Subdivisions[0].Names["zh-CN"]);
        }else{
            fmt.Printf("ip: %s \t 国家: %v \n", str, record.Country.Names["zh-CN"]);
        }
    }        
    {
        db, err := geoip2.Open("GeoLite2-City.mmdb")
        if err != nil {
                panic(err)
        }
        defer db.Close();

        // If you are using strings that may be invalid, check that ip is not nil
        str := "10.151.84.166";
        ip := net.ParseIP(str);
        record, err := db.City(ip)
        if err != nil {
                panic(err)
        }
        if len(record.Subdivisions)>0 {
            fmt.Printf("ip: %s \t 国家: %v \t 城市: %v \t 省份: %s\n", 
                str, record.Country.Names["zh-CN"], record.City.Names["zh-CN"],record.Subdivisions[0].Names["zh-CN"]);
        }else{
            fmt.Printf("ip: %s \t 国家: %v \n", str, record.Country.Names["zh-CN"]);
        }
    }                 
    
}