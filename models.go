package geoip

import (
	"encoding/json"
)

type IP struct {
	IP          string   `json:"ip"`
	CountryName string   `json:"country_name"`
	CountryCode string   `json:"country_code"`
	RegionName  string   `json:"region_name"`
	RegionCode  string   `json:"region_code"`
	City        string   `json:"city"`
	ZipCode     string   `json:"zipcode"`
	Timezone    Timezone `json:"timezone"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	MetroCode   int      `json:"metro_code"`
	ISP         string   `json:"isp"`
	Org         string   `json:"org"`
	As          string   `json:"as"`
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
type IPStackResult struct {
	*IPResult
	IP          string   `json:"ip"`
	CountryCode string   `json:"country_code"`
	CountryName string   `json:"country_name"`
	RegionCode  string   `json:"region_code"`
	RegionName  string   `json:"region_name"`
	City        string   `json:"city"`
	ZipCode     string   `json:"zip"`
	Timezone    Timezone `json:"time_zone"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	MetroCode   int      `json:"metro_code"`
}

type Timezone struct {
	ID          string `json:"id"`
	CurrentTime string `json:"current_time"`
	GMTOffset   int    `json:"gmt_offset"`
	Code        string `json:"code"`
	IsDST       bool   `json:"is_daylight_saving"`
}

// {
//   "ip": "134.201.250.155",
//   "hostname": "134.201.250.155",
//   "type": "ipv4",
//   "continent_code": "NA",
//   "continent_name": "North America",
//   "country_code": "US",
//   "country_name": "United States",
//   "region_code": "CA",
//   "region_name": "California",
//   "city": "Los Angeles",
//   "zip": "90013",
//   "latitude": 34.0453,
//   "longitude": -118.2413,
//   "location": {
//     "geoname_id": 5368361,
//     "capital": "Washington D.C.",
//     "languages": [
//         {
//           "code": "en",
//           "name": "English",
//           "native": "English"
//         }
//     ],
//     "country_flag": "https://assets.ipstack.com/images/assets/flags_svg/us.svg",
//     "country_flag_emoji": "🇺🇸",
//     "country_flag_emoji_unicode": "U+1F1FA U+1F1F8",
//     "calling_code": "1",
//     "is_eu": false
//   },
//   "time_zone": {
//     "id": "America/Los_Angeles",
//     "current_time": "2018-03-29T07:35:08-07:00",
//     "gmt_offset": -25200,
//     "code": "PDT",
//     "is_daylight_saving": true
//   },
//   "currency": {
//     "code": "USD",
//     "name": "US Dollar",
//     "plural": "US dollars",
//     "symbol": "$",
//     "symbol_native": "$"
//   },
//   "connection": {
//     "asn": 25876,
//     "isp": "Los Angeles Department of Water & Power"
//   },
//   "security": {
//     "is_proxy": false,
//     "proxy_type": null,
//     "is_crawler": false,
//     "crawler_name": null,
//     "crawler_type": null,
//     "is_tor": false,
//     "threat_level": "low",
//     "threat_types": null
//   }
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
