package geoip

import (
	"encoding/json"
)

type IP struct {
	IP          string  `json:"ip"`
	CountryName string  `json:"country_name"`
	CountryCode string  `json:"country_code"`
	RegionName  string  `json:"region_name"`
	RegionCode  string  `json:"region_code"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zipcode"`
	Timezone    string  `json:"timezone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

func (ip *IP) ToJSONString() string {
	data, err := json.Marshal(ip)
	if err != nil {
		return ""
	}
	return string(data)
}

type IPResult struct {
}

type DebugResult struct {
	RawResponse  []byte `json:"-"`
	RequestedURL string `json:"requested_url"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
}

// freegeoip.net IP
type FreeGeoIPResult struct {
	*IPResult
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	Timezone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

// {
//     "ip": "208.80.152.201",
//     "country_code": "US",
//     "country_name": "United States",
//     "region_code": "CA",
//     "region_name": "California",
//     "city": "San Francisco",
//     "zip_code": "94105",
//     "time_zone": "America/Los_Angeles",
//     "latitude": 37.7898,
//     "longitude": -122.3942,
//     "metro_code": 807
// }

// ip-api.com IP
// usage limits
// http://ip-api.com/docs/api:json#usage_limits
type IPApiResult struct {
	*IPResult
	IP          string  `json:"ip"`
	CountryCode string  `json:"countryCode"`
	CountryName string  `json:"country"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip"`
	Timezone    string  `json:"timezone"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

// {
//     "as": "AS14907 Wikimedia Foundation, Inc.",
//     "city": "San Francisco",
//     "country": "United States",
//     "countryCode": "US",
//     "isp": "Wikimedia Foundation, Inc.",
//     "lat": 37.7898,
//     "lon": -122.3942,
//     "org": "Wikimedia Foundation, Inc.",
//     "query": "208.80.152.201",
//     "region": "CA",
//     "regionName": "California",
//     "status": "success",
//     "timezone": "America/Los_Angeles",
//     "zip": "94105"
// }
