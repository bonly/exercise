package main

import (
    "encoding/json"
    "errors"
    "net/http"
    "net/url"
    "fmt"
)

const GEOCODE_API = "http://maps.googleapis.com/maps/api/geocode/json"

type Geocode struct {
    Status  string `json:"status"`
    Results []Result `json:"results"`
}

type Result struct {
    AddressComponents []AddressComponent `json:"address_components"`
    FormattedAddress  string             `json:"formatted_address"`
    Geometry          Geometry           `json:"geometry"`
    Types             []string           `json:"types"`
}

type AddressComponent struct {
    LongName  string   `json:"long_name"`
    ShortName string   `json:"short_name"`
    Types     []string `json:types`
}

type Geometry struct {
    Bounds       Bounds     `json:"bounds"`
    Location     Coordinate `json:"location"`
    LocationType string     `json:"location_type"`
    Viewport     Bounds     `json:"viewport"`
}

type Bounds struct {
    Northeast Coordinate `json:"northwest"`
    Southwest Coordinate `json:"southwest"`
}

type Coordinate struct {
    Lat float64 `json:"lat"`
    Lng float64 `json:"lng"`
}

func fecthParam(location string, lat, lng float64)(ret string){
    qs, _ := url.ParseQuery("");
    qs.Add("sensor", "false");
    qs.Add("language", "zh_CN");
    if location != "" {
        qs.Add("address", location);
    }
    if lat != 0 || lng != 0{
        latlng := fmt.Sprintf("%f,%f", lat, lng);
        qs.Add("latlng", latlng);
    }

    ret = GEOCODE_API + "?" + qs.Encode();
    return ret;
}

func fetchGeocode(city string, lat, lng float64) (*Geocode, error) {
    strUrl := fecthParam(city, lat, lng);
    var geo Geocode

    resp, err := http.Get(strUrl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
        return nil, err
    }
    //fmt.Println(geo);

    return &geo, nil
}

func GetLocation(city string) (float64, float64, error) {
    geo, err := fetchGeocode(city,0,0)
    if err != nil {
        return 0, 0, err
    }

    if len(geo.Results) == 0 {
        return 0, 0, errors.New("Unkown location")
    }

    return geo.Results[0].Geometry.Location.Lat, geo.Results[0].Geometry.Location.Lng, nil
}

func GetCity(lat, lng float64) (loc string, err error){
    geo, err := fetchGeocode("",lat,lng)
    if err != nil {
        return "",err;
    }

    if len(geo.Results) == 0 {
        return "",errors.New("Unkown location");
    } 
    var country,province,city string;
    for _, ac := range geo.Results[0].AddressComponents {
        //fmt.Println(i, "\t", ac.Types, "\t", ac.LongName);
        if ac.Types[0] == "country" {
           country = ac.LongName;
        }
        if ac.Types[0] == "administrative_area_level_1"{
            province = ac.LongName;
        }
        if ac.Types[0] == "locality"{
            city = ac.LongName;
        }        
    }
    ret := country+","+province+","+city;

    return ret, nil;
}

func main(){
    {
        geo, err := fetchGeocode("Berlin",0,0)
        if err != nil {
            fmt.Println("%v", err)
        }

        if len(geo.Results) == 0 {
            fmt.Println("Geo result are empty")
        }

        if geo.Results[0].Geometry.Location.Lat != 52.5191710 {
            fmt.Println("Expected latidute of 52.5191710 got %v", geo.Results[0].Geometry.Location.Lat)
        }
    }

    {
        lat, lng, err := GetLocation("Berlin")
        if err != nil {
            fmt.Println("%v", err)
        }

        if lat != 52.5191710 {
            fmt.Println("Expected latidute of 52.5191710 got %v", lat)
        }

        if lng != 13.40609120 {
            fmt.Println("Expected longtitude of 13.40609120 got %v", lng)
        }   
    }
    {
        ret, err := GetCity(23.129206, 113.277201);
        if err != nil {
            fmt.Println("%v", err)
        }
        fmt.Println(ret);
    }    
    {
        ret, err := GetCity(30.540859,120.7768689);
        if err != nil {
            fmt.Println("%v", err)
        }
        fmt.Println(ret);
    }       
    {
        ret, err := GetCity(22.535586000000002,113.921386842857);
        if err != nil {
            fmt.Println("%v", err)
        }
        fmt.Println(ret);
    }      
}