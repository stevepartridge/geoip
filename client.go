package geoip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	Debug      bool
	LastLookup *IP
}

func New() *Client {
	return &Client{}
}

// Lookup takes a string based IPv4 address and queries for location and ISP info.
func (self *Client) Lookup(ipQuery string) (IP, error) {

	ip := IP{}
	var err error

	funcs := []func(ip *IP){

		// freegeoip.net
		func(ip *IP) {
			freeGeoIP := FreeGeoIPResult{}
			var debug DebugResult
			debug, err = self.call(
				"http://freegeoip.net/json/"+ipQuery,
				&freeGeoIP,
			)

			if err != nil {
				fmt.Println("error", err)
				return
			}

			ip.IP = freeGeoIP.IP
			ip.CountryName = freeGeoIP.CountryName
			ip.CountryCode = freeGeoIP.CountryCode
			ip.RegionName = freeGeoIP.RegionName
			ip.RegionCode = freeGeoIP.RegionCode
			ip.City = freeGeoIP.City
			ip.ZipCode = freeGeoIP.ZipCode
			ip.Timezone = freeGeoIP.Timezone
			ip.Latitude = freeGeoIP.Latitude
			ip.Longitude = freeGeoIP.Longitude
			ip.MetroCode = freeGeoIP.MetroCode

			if self.Debug {
				fmt.Println("debug freegeoip", debug, freeGeoIP)
			}
		},

		// ip-api.com
		func(ip *IP) {
			ipApi := IPApiResult{}
			var debug DebugResult
			debug, err = self.call(
				"http://ip-api.com/json/"+ipQuery,
				&ipApi,
			)

			if err != nil {
				fmt.Println("error", err)
				return
			}

			ip.ISP = ipApi.ISP
			ip.As = ipApi.As
			ip.Org = ipApi.Org

			if self.Debug {
				fmt.Println("debug ip-api", debug, ipApi)
			}
		},
	}

	wg := new(sync.WaitGroup)

	wg.Add(len(funcs))

	for _, fn := range funcs {
		go func(f func(ip *IP)) {
			f(&ip)
			wg.Done()
		}(fn)
	}

	wg.Wait()

	if self.Debug {
		fmt.Println("IP", ip, err)
	}

	return ip, err

}

// Call uses basic (GET) method to make a request to the API
func (self *Client) call(apiUrl string, result interface{}) (DebugResult, error) {

	timeout := time.Duration(10 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(apiUrl)

	if err != nil {
		return DebugResult{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return DebugResult{}, err
	}

	debug := DebugResult{
		RequestedURL: apiUrl,
		Status:       resp.Status,
		StatusCode:   resp.StatusCode,
		RawResponse:  body,
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return debug, err
	}

	return debug, err
}
