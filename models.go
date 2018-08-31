package geoip

import (
	"encoding/json"
)

type IP struct {
	IP          string     `json:"ip"`
	CountryName string     `json:"country_name"`
	CountryCode string     `json:"country_code"`
	RegionName  string     `json:"region_name"`
	RegionCode  string     `json:"region_code"`
	City        string     `json:"city"`
	ZipCode     string     `json:"zipcode"`
	Timezone    Timezone   `json:"timezone"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	MetroCode   int        `json:"metro_code"`
	ISP         string     `json:"isp"`
	Org         string     `json:"org"`
	As          string     `json:"as"`
	Location    Location   `json:"location"`
	Currency    Currency   `json:"currency"`
	Connection  Connection `json:"connection"`
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
	IPResult
	IP            string   `json:"ip"`
	Hostname      string   `json:"hostname"`
	Type          string   `json:"type"`
	ContinentCode string   `json:"continent_code"`
	ContinentName string   `json:"continent_name"`
	CountryCode   string   `json:"country_code"`
	CountryName   string   `json:"country_name"`
	RegionCode    string   `json:"region_code"`
	RegionName    string   `json:"region_name"`
	City          string   `json:"city"`
	ZipCode       string   `json:"zip"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	Location      Location `json:"location"`
	Timezone      Timezone `json:"time_zone"`
	Currency      Currency `json:"currency"`
	MetroCode     int      `json:"metro_code"`
}

type Timezone struct {
	ID          string `json:"id"`
	CurrentTime string `json:"current_time"`
	GMTOffset   int    `json:"gmt_offset"`
	Code        string `json:"code"`
	IsDST       bool   `json:"is_daylight_saving"`
}

//   "time_zone": {
//     "id": "America/Los_Angeles",
//     "current_time": "2018-03-29T07:35:08-07:00",
//     "gmt_offset": -25200,
//     "code": "PDT",
//     "is_daylight_saving": true
//   },

type Location struct {
	GeonameId               int                      `json:"geoname_id"`
	Capital                 string                   `json:"capital"`
	Languages               []map[string]interface{} `json:"languages"`
	CountryFlag             string                   `json:"country_flag"`
	CountryFlagEmoji        string                   `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string                   `json:"country_flag_emoji_unicode"`
	CallingCode             string                   `json:"calling_code"`
	IsEU                    bool                     `json:"is_eu"`
}

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

type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Plural       string `json:"plural"`
	Symbol       string `json:"symbol"`
	SymbolNative string `json:"symbol_native"`
}

//   "currency": {
//     "code": "USD",
//     "name": "US Dollar",
//     "plural": "US dollars",
//     "symbol": "$",
//     "symbol_native": "$"
//   },

type Connection struct {
	ASN int    `json:"asn"`
	ISP string `json:"isp"`
}

//   "connection": {
//     "asn": 25876,
//     "isp": "Los Angeles Department of Water & Power"
//   },

type Security struct {
	IsProxy bool `json:"is_proxy"`
	// ProxyType
	IsCrawler bool `json:"is_crawler"`
	// CrawlerName
	// CrawlerType
	IsTOR       bool   `json:"is_tor"`
	ThreatLevel string `json:"threat_level"`
	// ThreatTypes string `json:"threat_types"`
}

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
// }

// ip-api.com IP
// usage limits
// http://ip-api.com/docs/api:json#usage_limits
type IPApiResult struct {
	IPResult
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
