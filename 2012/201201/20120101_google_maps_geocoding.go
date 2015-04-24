package geo

import (
	"../rest"
	"encoding/json"
	//"log"
	"net/url"
)

type LatLng struct {
	Lat float64 "lat"
	Lng float64 "lng"
}

type GoogleMapsResponse struct {
	Results []GoogleMapsResults
	Status  string
}

type GoogleMapsResults struct {
	Geometry GoogleMapsGeometry "geometry"
	Types    []string           "types"
}

type GoogleMapsGeometry struct {
	Location LatLng "location"
}

func GetLatLng(addr string, city string, state string) (latlong LatLng, err error) {
	latlong.Lat = 0
	latlong.Lng = 0
	urlstr := "http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=" + url.QueryEscape(addr+" "+city+", "+state)
	response, err := rest.Get(urlstr)
	if err != nil {
		return latlong, err
	}
	gmapsResp := GoogleMapsResponse{}
	err = json.Unmarshal(response, &gmapsResp)
	if err != nil {
		return latlong, err
	}
	if gmapsResp.Status == "OK" && len(gmapsResp.Results) > 0 {
		latlong = gmapsResp.Results[0].Geometry.Location
	}
	return latlong, nil
}

