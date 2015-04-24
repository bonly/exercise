package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	geocodingURL = "http://maps.googleapis.com/maps/api/geocode/json"
)
//bool, string, int, int64可在声明中指定类型如： Result `json: "result,int"`
//字段中"-"，那么这个字段不会输出到JSON
//omitempty:如果该字段值为空，就不会输出到JSON串中

type GeocodingResult struct {
	Results []*GeocodingAddress `json:"results,omitempty"`  
}

type GeocodingAddress struct {
	FormattedAddress string            `json:"formatted_address,omitempty"`
	Geometry         GeocodingGeometry `json:"geometry,omitempty"`
}

type GeocodingGeometry struct {
	Location GeometryLocation `json:"location,omitempty"`
}

type GeometryLocation struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

func (l *GeometryLocation) String() string {
	return fmt.Sprintf("%v, %v", l.Lat, l.Lng)
}

func Geocoding(location string, lat, lng float64) (*GeocodingResult, error) {
	qs, err := url.ParseQuery("")
	Check(err)
	if location != "" {
		qs.Add("address", location)
	}
    if lat != 0 || lng != 0{
        latlng := fmt.Sprintf("%f,%f", lat, lng);
        qs.Add("latlng", latlng);
    }
	qs.Add("sensor", "false")
	qs.Add("language", "zh_CN")

	u := geocodingURL + "?" + qs.Encode()

	resp, err := http.Get(u)
	Check(err)

	r := new(GeocodingResult)
	err = json.NewDecoder(resp.Body).Decode(r)
	return r, err
}
 
 func main(){
 	{
	 	r, _ := Geocoding("广州丽江花园", 0, 0);
	    fmt.Printf("%s\n", r.Results[0].FormattedAddress);
	    fmt.Printf("%v\n", r.Results[0].Geometry.Location);	
	}
	{
	 	r, _ := Geocoding("广州", 23.035441, 113.300614);
	    fmt.Printf("%s\n", r.Results[0].FormattedAddress);
	    fmt.Printf("%v\n", r.Results[0].Geometry.Location);			
	}
	{
	 	r, _ := Geocoding("", 23.035441, 113.300614);
	    fmt.Printf("%s\n", r.Results[0].FormattedAddress);
	    fmt.Printf("%v\n", r.Results[0].Geometry.Location);			
	}	
 }

 func Check(err error){
 	if err != nil{
 		fmt.Println(err);
 	}
 }